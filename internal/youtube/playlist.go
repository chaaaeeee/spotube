package youtube

import (
	"io"
	"net/http"
	"fmt"

	_ "github.com/chaaaeeee/spotube/config")

func SearchVideos(client *http.Client) {
	searchURL := "https://www.googleapis.com/youtube/v3/search?part=snippet&q=Not+Like+Us+Kendrick+Lamar&key=AIzaSyBTuHTwvKOeHkReGGHpxEnc9t-1gPE24H0"
	req, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)	
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
