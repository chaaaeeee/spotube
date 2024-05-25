package main

import (
	"github.com/chaaaeeee/spotube/internal/spotify"
)

func main() {
	client := spotify.NewClient()
	client.GetPlaylistTracks()
}
