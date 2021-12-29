package main

// make a config file for tokens,environmnt
// show me the tweets from a particular person in the ..
// show me the tweets for a particular hashtag in a time frame one month
// save the topics

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var creds Credentials

func main() {

	if os.Getenv("API_ENVIRONMENT") == "twitter" {
		fmt.Println("we are in twitter environment")

		creds = Credentials{
			AccessToken:       os.Getenv("ACCESS_TOKEN"),
			AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
			ConsumerKey:       os.Getenv("CONSUMER_KEY"),
			ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
		}
		client, err := TwitterCredentialsCheck(&creds)
		if err != nil {
			log.Println("Error getting Twitter Client")
			log.Println(err)
			return
		}

		searchCmd := flag.NewFlagSet("search", flag.ExitOnError)
		searchString := searchCmd.String("topic", "", "What to search?")

		showCmd := flag.NewFlagSet("show", flag.ExitOnError)
		showId := showCmd.Int64("id", 0, "What's the id of the tweet?")

		flag.NewFlagSet("trends", flag.ExitOnError)

		followersCmd := flag.NewFlagSet("followers", flag.ExitOnError)
		followersId := followersCmd.Int64("id", 0, "What's the user id?")

		if len(os.Args) < 2 {
			fmt.Println("expected 'search' or 'show' subcommands")
			os.Exit(1)
		}

		//look at the 2nd argument's value
		switch os.Args[1] {
		case "search":
			handleSearch(client, searchCmd, searchString)
		case "show":
			handleShow(client, showCmd, showId)
		case "trends":
			trends(client)
		case "followers":
			followers(client, followersCmd, followersId)
		default:
			fmt.Println("we are currently in twitter environment")
			fmt.Println("We only have 'search' and 'show' subcommands ")
		}
	} else if os.Getenv("API_ENVIRONMENT") == "linkedin" {
		fmt.Println("we are in linkedin environment")
	} else {
		fmt.Println("We only have 'twitter' and 'linkedin' environments ")
		fmt.Println("RUN: setx API_ENVIRONMENT 'linkedin' ")
		fmt.Println("RUN: setx API_ENVIRONMENT 'twitter' ")
	}

}
