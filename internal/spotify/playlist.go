package spotify 

import (
	"net/http"
	"encoding/json"
	"fmt"
	"io"
)

type Playlist struct {
	Tracks Tracks `json:"tracks"`
}

type Tracks struct {
	Items []Item `json:"items"`
}

type Item struct {
	Track Track `json:"track"`
}

type Track struct {
	Name string `json:"name"`
	Artists []Artist `json:"artists"`
}

type Artist struct {
	Name string `json:"name"`
}

func getTracks(client *http.Client, token Token) Playlist {
	playlistUrl := apiUrl + "/playlists/6eYWS7Wy5x46dqND2HYP9S?fields=tracks.items%28track%28name%2C+artists%29%29"

	req, err := http.NewRequest("GET", playlistUrl, nil)
	if err != nil {
		panic(err)
	}	

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var playlist Playlist
	err = json.Unmarshal(body, &playlist)
	if err != nil {
		panic(err)
	}

	return playlist
}

func printTracks(playlist Playlist) {
	for _, tracks := range playlist.Tracks.Items {
		fmt.Println("Title :", tracks.Track.Name)
		for _, artists := range tracks.Track.Artists {
			fmt.Println("Artist :", artists.Name)
		}
	}
}
