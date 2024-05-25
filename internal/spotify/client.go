package spotify 

import (
	"strings"
	"net/http"
	"encoding/base64"
	"net/url"
	"io"
	"encoding/json"

	"github.com/zmb3/spotify/v2"
	"github.com/chaaaeeee/spotube"
	"golang.org/x/oauth2/clientcredentials"
)

type Client struct {
	spotifyClient *spotify.Client
}

func NewClient(clientID string, clientSecret string) *Client {
	clientConfig := clientcredentials.Config{
		ClientID : config.SpotifyClientId,
		ClientSecret : 
	}

	spotifyClient := spotify.New()

	return &Client{spotifyClient : spotifyClient}
}

func GetAccessToken(client *http.Client) Token {
	var token Token
	authValue := "Basic " + base64.StdEncoding.EncodeToString([]byte(config.SpotifyClientId + ":" + config.SpotifyClientSecret))

	form := url.Values{}
	form.Add("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", config.SpotifyAuthBaseURL, strings.NewReader(form.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", authValue)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	
	err = json.Unmarshal(body, &token)
	if err != nil {
		panic(err)
	}

	return token
}
