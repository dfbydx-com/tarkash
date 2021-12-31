package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	// "golang.org/x/oauth2"
	// "golang.org/x/oauth2/clientcredentials"
)

type tweet struct {
	ID   int64
	Text string
	//User *twitter.User
	CreatedAt string
	Entities  *twitter.Entities
}

type Credentials struct {
	ConsumerKey       string //api key
	ConsumerSecret    string //api secret
	AccessToken       string
	AccessTokenSecret string
}

func TwitterCredentialsCheck(creds *Credentials) (*twitter.Client, error) {
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}
	// we can retrieve the user and verify if the credentials
	// we have used successfully allow us to log in!
	_, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}
	fmt.Println("got the twitter client...")
	//log.Printf("User's ACCOUNT:\n%+v\n", user)
	return client, nil
}

func search(client *twitter.Client, query string) []tweet {
	search, _, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query:      query,
		Count:      60,
		ResultType: "popular",
		Since:      time.Now().AddDate(0, -1, 0).Format("2006-01-02"),
	})
	if err != nil {
		log.Print(err)
	}
	search1 := search.Statuses
	var search2 []tweet
	var x tweet
	for _, t := range search1 {
		x.ID = t.ID
		x.Text = t.Text
		//x.User = t.User
		x.CreatedAt = t.CreatedAt
		x.Entities = t.Entities
		search2 = append(search2, x)
	}
	return search2
}
func handleSearch(client *twitter.Client, searchCmd *flag.FlagSet, all *bool, text *string) {
	searchCmd.Parse(os.Args[2:])
	if *text == "" && !*all {
		fmt.Print("topic is required to search or specify --all")
		searchCmd.PrintDefaults()
		os.Exit(1)
	} else {
		var searchResult []tweet
		if *all {
			fmt.Println("____________searching for all added topics______________\n")
			topics := getTopics()
			for _, x := range topics {
				var boundary tweet
				boundary.Text = "____________________searching for '" + x.Text + "' now____________________"
				searchResult = append(searchResult, boundary)
				searchResult = append(searchResult, search(client, x.Text)...)
			}
			print(searchResult)
		} else {
			fmt.Printf("____________searching for %v______________\n", *text)
			searchResult = append(searchResult, search(client, *text)...)
			print(searchResult)
		}
		fmt.Println("Search results are stored in 'print.json' file")
	}
}

// func send(client *twitter.Client) {
// 	fmt.Println("____________sending tweets______________")
// 	tweet, resp, err := client.Statuses.Update("This tweet is from a bot I am building using golang!", nil)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	log.Printf("%+v\n", resp)
// 	log.Printf("%+v\n", tweet)
// }

func trends(client *twitter.Client) {
	l, _, err := client.Trends.Available()
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("%v", l)
}

func followers(client *twitter.Client, followersCmd *flag.FlagSet, id *int64) {
	followersCmd.Parse(os.Args[2:])
	if *id == 0 {
		fmt.Print("user id is required to show the followers")
		followersCmd.PrintDefaults()
		os.Exit(1)
	} else {
		followerList, _, err := client.Followers.List(&twitter.FollowerListParams{
			UserID: *id,
		})
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("%v", followerList)
	}
}

func userTimeline(client *twitter.Client, userTimelineScreenName *string) []tweet {
	includeRetweets := false
	timeline, _, err := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		ScreenName: *userTimelineScreenName,
		//TrimUser:        flag.Bool("faltu", true, "faltu"),
		IncludeRetweets: &includeRetweets,
		Count:           60,
	})
	if err != nil {
		log.Print(err)
	}
	var timeline1 []tweet
	var x tweet
	for _, t := range timeline {
		x.ID = t.ID
		x.Text = t.Text
		x.CreatedAt = t.CreatedAt
		x.Entities = t.Entities
		timeline1 = append(timeline1, x)
	}
	return timeline1
}
func handleUserTimeline(client *twitter.Client, userTimelineCmd *flag.FlagSet, all *bool, userTimelineScreenName *string) {
	userTimelineCmd.Parse(os.Args[2:])
	if *userTimelineScreenName == "" && !*all {
		fmt.Print("Screen Name is required to see the timeline of the user or specify --all")
		userTimelineCmd.PrintDefaults()
		os.Exit(1)
	} else {
		var timelineResult []tweet
		if *all {
			fmt.Println("____________searching for all added users______________\n")
			users := getUsers()
			for _, x := range users {
				var boundary tweet
				boundary.Text = "____________________searching for '@" + x.Screen_Name + "' now____________________"
				timelineResult = append(timelineResult, boundary)
				timelineResult = append(timelineResult, userTimeline(client, &x.Screen_Name)...)
			}
			print(timelineResult)
		} else {
			fmt.Printf("____________searching for @%v______________\n", *userTimelineScreenName)
			timelineResult = append(timelineResult, userTimeline(client, userTimelineScreenName)...)
			print(timelineResult)
		}
		fmt.Println("Results are stored in 'print.json' file")
	}
}
