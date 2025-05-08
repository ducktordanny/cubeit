package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ducktordanny/cubeit/backend/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (store *Store) RegisterOrUpdateUser(wcaUser types.WCAUser) error {
	fmt.Printf("Everything is fine. User %s\n", wcaUser.Name)
	return nil
}

func (store *Store) GetWCAUser(accessToken string) (types.WCAUser, error) {
	meURL := "https://www.worldcubeassociation.org/api/v0/me"
	req, err := http.NewRequest("GET", meURL, nil)
	if err != nil {
		return types.WCAUser{}, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return types.WCAUser{}, fmt.Errorf("request failed: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return types.WCAUser{}, fmt.Errorf(
				"WCA user fetch failed (%d); also failed to read body: %w", res.StatusCode, err,
			)
		}
		return types.WCAUser{}, fmt.Errorf(
			"WCA user fetch failed (%d): %s", res.StatusCode, body,
		)
	}

	var wcaMe types.WCAMe
	if err := json.NewDecoder(res.Body).Decode(&wcaMe); err != nil {
		return types.WCAUser{}, fmt.Errorf("failed to decode user response: %w", err)
	}
	return wcaMe.Me, nil
}
