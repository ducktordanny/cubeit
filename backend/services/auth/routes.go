package auth

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/ducktordanny/cubeit/backend/configs"
	"github.com/ducktordanny/cubeit/backend/types"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (handler *Handler) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("auth/login", handleLogin)
	router.GET("oauth/callback", handler.handleOAuthCallback)
}

func handleLogin(context *gin.Context) {
	state, err := generateState(32)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate state"})
		return
	}

	context.SetCookie("oauth_state", state, 300, "/", "", configs.Envs.Production, true)

	authURL, err := url.Parse("https://www.worldcubeassociation.org/oauth/authorize")
	if err != nil {
		log.Fatalf("Invalid base URL: %v", err)
	}

	params := url.Values{}
	params.Add("client_id", configs.Envs.ClientID)
	params.Add("redirect_uri", configs.Envs.RedirectURI)
	params.Add("response_type", "code")
	params.Add("scope", "public email")
	params.Add("state", state)

	authURL.RawQuery = params.Encode()
	finalURL := authURL.String()

	context.Redirect(http.StatusFound, finalURL)
}

func (handler *Handler) handleOAuthCallback(context *gin.Context) {
	code := context.Query("code")
	state := context.Query("state")
	expectedState, err := context.Cookie("oauth_state")
	if err != nil || state != expectedState {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing state"})
		return
	}

	authInfo, err := getAuthInfo(code)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to exchange code: " + err.Error(),
		})
		return
	}

	wcaUser, err := handler.store.GetWCAUser(authInfo.AccessToken)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get WCA user: " + err.Error(),
		})
		return
	}

	user, err := handler.store.RegisterOrUpdateUser(wcaUser)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to register/update user: " + err.Error(),
		})
		return
	}
	context.IndentedJSON(http.StatusOK, user)
}

func generateState(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(bytes), nil
}

func getAuthInfo(code string) (types.AuthorizationInfo, error) {
	form := url.Values{
		"grant_type":    {"authorization_code"},
		"client_id":     {configs.Envs.ClientID},
		"client_secret": {configs.Envs.ClientSecret},
		"code":          {code},
		"redirect_uri":  {configs.Envs.RedirectURI},
	}

	tokenURL := "https://www.worldcubeassociation.org/oauth/token"
	res, err := http.PostForm(tokenURL, form)
	if err != nil {
		return types.AuthorizationInfo{}, fmt.Errorf("Request failed: %w", err)
	}
	defer res.Body.Close()

	var auth types.AuthorizationInfo
	if err := json.NewDecoder(res.Body).Decode(&auth); err != nil {
		return types.AuthorizationInfo{}, fmt.Errorf("Decode failed: %w", err)
	}

	return auth, nil
}
