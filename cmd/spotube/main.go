package main

import (
	"fmt"
	"net/http"

	"github.com/chaaaeeee/spotube/internal/spotify"
)

func main() {
	client := &http.Client{}

	token := spotify.GetAccessToken(client)
	fmt.Println(token)
	playlist := spotify.GetTracks(client, token)

	spotify.PrintTracks(playlist)
}
