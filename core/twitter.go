package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/auliman-a/DiscordBot/entities"
)

var Token string

func GetTweet(username string) string {
	url := "https://api.twitter.com/1.1/statuses/user_timeline.json?screen_name=" + username + "&count=1"

	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + Token

	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	jsonStr := string([]byte(body))

	tweet := []entities.TwitterFeed{}

	json.Unmarshal([]byte(jsonStr), &tweet)

	fmt.Println(string([]byte(body)))

	return tweet[0].Text
}
