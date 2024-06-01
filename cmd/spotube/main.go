package main

import (
	"github.com/chaaaeeee/spotube/internal/youtube"
	"net/http"
)

func main() {
	http.HandleFunc("/google/login", youtube.GoogleLogin)
	http.HandleFunc("/google/callback", youtube.GoogleCallback)

	http.ListenAndServe(":8080", nil)
}
