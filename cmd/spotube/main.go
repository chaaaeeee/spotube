package main

import (
	"fmt"
	"net/http"

	"github.com/chaaaeeee/spotube/internal"
)

func main() {
	client := &http.Client{}

	token := GetAccessToken(client)
	fmt.Println(token)
	playlist := getTracks(client, token)

	printTracks(playlist)
}
