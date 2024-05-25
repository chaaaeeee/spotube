package main

import (
	"net/http"

	"github.com/chaaaeeee/spotube/internal/spotify"
	"github.com/chaaaeeee/spotube/internal/youtube"
)

func main() {
	client := &http.Client{}
	token := spotify.GetAccessToken(client)
	playlist := spotify.GetTracks(client, token)

	spotify.PrintTracks(playlist)

	youtube.SearchVideos(client)
}
