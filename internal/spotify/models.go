package spotify

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
