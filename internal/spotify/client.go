package spotify

import (
	"context"

	"github.com/chaaaeeee/spotube/config"
	spo "github.com/zmb3/spotify/v2"
	"golang.org/x/oauth2/clientcredentials"
)

type Client struct {
	spotifyClient *spo.Client
}

func NewClient() *Client {
	clientConfig := clientcredentials.Config{
		ClientID:     config.SpotifyClientId,
		ClientSecret: config.SpotifyClientSecret,
		TokenURL:     "https://accounts.spotify.com/api/token",
	}

	httpClient := clientConfig.Client(context.Background())

	spotifyClient := spo.New(httpClient)

	return &Client{spotifyClient: spotifyClient}
}
