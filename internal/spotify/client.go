package spotify

import (
	"context"
	"fmt"
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
		TokenURL: "https://accounts.spotify.com/api/token",
	}

	httpClient := clientConfig.Client(context.Background())

	spotifyClient := spo.New(httpClient)
	
	client := &Client{spotifyClient: spotifyClient}

	return client
}

func (c *Client) GetPlaylistTracks() {
	playlistItems, err := c.spotifyClient.GetPlaylistItems(context.Background(), "0E8fe06Z5G4xWKqwcSCuIC")
	if err != nil {
		panic(err)
	}

	for _, music := range playlistItems.Items {
		fmt.Println(music.Track.Track.SimpleTrack.Name)
	}
}

