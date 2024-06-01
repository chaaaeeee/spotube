package youtube

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"net/http"
	"os"
)

type OAuth2Config struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
}

type Service struct {
	youtubeService *youtube.Service
}

func NewService(r *http.Request) (*Service, error) {
	cookie, err := r.Cookie("oauth2token")
	if err != nil {
		return nil, fmt.Errorf("error retrieving cookie: %w", err)
	}

	var token oauth2.Token
	err = json.Unmarshal([]byte(cookie.Value), &token)
	if err != nil {
		return nil, fmt.Errorf("error decoding token: %w", err)
	}

	googleConfig := SetupConfig()
	client := googleConfig.Client(context.Background(), &token)

	youtubeService, err := youtube.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("error creating YouTube service: %w", err)
	}

	return &Service{youtubeService: youtubeService}, nil
}

func SetupConfig() *oauth2.Config {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	conf := &oauth2.Config{
		ClientID:     os.Getenv("ClientID"),
		ClientSecret: os.Getenv("ClientSecret"),
		RedirectURL:  os.Getenv("RedirectURL"),
		Scopes: []string{
			"https://www.googleapis.com/auth/youtube",
			"https://www.googleapis.com/auth/youtube.force-ssl",
		},
		Endpoint: google.Endpoint,
	}

	return conf
}

func GoogleLogin(w http.ResponseWriter, r *http.Request) {
	googleConfig := SetupConfig()
	url := googleConfig.AuthCodeURL("random")

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func GoogleCallback(w http.ResponseWriter, r *http.Request) {
	state := r.URL.Query()["state"][0]
	if state != "random" {
		fmt.Println("States don't match")
		panic("error")
	}
	code := r.URL.Query()["code"][0]

	googleConfig := SetupConfig()

	token, err := googleConfig.Exchange(context.Background(), code)
	if err != nil {
		fmt.Println("Error retrieving token")
		panic(err)
	}

	tokenJSON, err := json.Marshal(token)
	if err != nil {
		panic(err)
	}

	cookie := http.Cookie{
		Name:     "oauth2token",
		Value:    string(tokenJSON),
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)
}
