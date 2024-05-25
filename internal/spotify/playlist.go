package spotify

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/chaaaeeee/spotube/config"
)

func GetTracks(client *http.Client, token Token) PlaylistData {
	playlistUrl := config.SpotifyAPIBaseURL + "/playlists/3WW6fjWlVL7yzDysqyFoOw?fields=tracks.items%28track%28name%2C+artists%29%29"

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

	var playlist PlaylistData
	err = json.Unmarshal(body, &playlist)
	if err != nil {
		panic(err)
	}

	return playlist
}

func PrintTracks(playlistData PlaylistData) []string {
	var playlist []string
	var music string
	for _, tracks := range playlistData.Tracks.Items {
		for _, artists := range tracks.Track.Artists {
			music = fmt.Sprintf("%s - %s", tracks.Track.Name, artists.Name)
			break
		}

		playlist = append(playlist, music)
	}

	for _, music := range playlist {
		fmt.Println(music)
	}

	return playlist
}
