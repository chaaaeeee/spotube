package spotify 

import (
	"strings"
	"net/http"
	"encoding/base64"
	"net/url"
	"io"
	"encoding/json"
)

type Token struct {
	AccessToken string `json:"access_token"`
	Type string `json:"token_type"`
	Exp int `json:"expires_in"`
}

func GetAccessToken(client *http.Client) Token {
	var token Token
	authValue := "Basic " + base64.StdEncoding.EncodeToString([]byte(ClientId + ":" + ClientSecret))

	form := url.Values{}
	form.Add("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", authUrl, strings.NewReader(form.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", authValue)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	
	err = json.Unmarshal(body, &token)
	if err != nil {
		panic(err)
	}

	return token
}
