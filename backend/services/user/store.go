package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ducktordanny/cubeit/backend/configs"
	"github.com/ducktordanny/cubeit/backend/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (store *Store) RegisterOrUpdateUser(wcaUser types.WCAUser) (types.User, error) {
	_, err := store.db.Exec(`
		INSERT INTO "user" ("id", "wcaId", "name", "email", "gender", "countryISO", "avatarURL")
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT ("id") DO UPDATE
		SET "wcaId" = EXCLUDED."wcaId",
		    "name" = EXCLUDED."name",
		    "email" = EXCLUDED."email",
		    "avatarURL" = EXCLUDED."avatarURL"`,
		wcaUser.Id, wcaUser.WcaId, wcaUser.Name, wcaUser.Email,
		wcaUser.Gender, wcaUser.CountryIso2, wcaUser.Avatar.Url,
	)
	if err != nil {
		return types.User{}, err
	}
	var user types.User
	err = store.db.QueryRow(`SELECT * FROM "user" WHERE "id" = $1`, wcaUser.Id).Scan(
		&user.Id, &user.WcaId, &user.Name, &user.Email, &user.Gender, &user.Bio,
		&user.CountryISO, &user.AvatarURL, &user.Role, &user.CreatedAt,
	)
	if err != nil {
		return types.User{}, err
	}
	return user, nil
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
	if configs.Envs.Production != true {
		fmt.Printf("AccessToken of %s: %s\n", wcaMe.Me.Name, accessToken)
	}
	return wcaMe.Me, nil
}
