package domain

type Token struct {
	AccessToken string `json:"access_token"`
	Type string `json:"token_type"`
	Exp int `json:"expires_in"`
}

type Playlist struct {
	Title []string
}

type PlaylistData struct {
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
