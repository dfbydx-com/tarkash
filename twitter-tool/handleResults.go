package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/dghubble/go-twitter/twitter"
)

type tweet struct {
	ID   int64
	Text string
	User *twitter.User
}

func saveResults(tweets []tweet, whattodo string) {
	tweetBytes, err := json.Marshal(tweets)
	if err != nil {
		panic(err)
	}
	if whattodo == "print" {
		err = ioutil.WriteFile("./print.json", tweetBytes, 0644)
		if err != nil {
			panic(err)
		}
	}

}
