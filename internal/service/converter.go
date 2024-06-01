package converter

import (
	"context"
	"fmt"

	"github.com/zmb3/spotify/v2"
	"google.golang.org/api/youtube/v3"
)

type Converter struct {
	spotifyClient  *spotify.Client
	youtubeService *youtube.Service
}

type ConverterStore interface {
	GetPlaylistPrivacy() bool
	GetPlaylistTracks() ([]string, error)
}

func (c *Converter) GetPlaylistTracks() ([]string, error) {
	// handle if playlist if private
	playlistItems, err := c.spotifyClient.GetPlaylistItems(context.Background(), "0E8fe06Z5G4xWKqwcSCuIC")
	if err != nil {
		return nil, err
	}

	var tracklist []string
	var track string
	for _, tracks := range playlistItems.Items {
		fmt.Println(tracks.Track.Track.SimpleTrack.Name)
		for _, artist := range tracks.Track.Track.SimpleTrack.Artists {
			track = fmt.Sprintf("%s - %s", tracks.Track.Track.SimpleTrack.Name, artist.Name)
		}

		tracklist = append(tracklist, track)
	}

	for _, item := range tracklist {
		fmt.Println(item)
	}

	return tracklist, nil
}
