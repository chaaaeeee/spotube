package youtube

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"net/http"
)

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
	conf := &oauth2.Config{
		ClientID:     "695816638410-57e60c4tiubt4hnc0obvf4fdstmi38br.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  "http://localhost:8080/google/callback",
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
