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

func search(client *twitter.Client, query string) {
	fmt.Printf("____________searching for %v______________\n", query)
	search, _, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query:      query,
		Count:      60,
		ResultType: "popular",
		Since:      time.Now().AddDate(0, -1, 0).Format("2006-01-02"),
	})
	if err != nil {
		log.Print(err)
	}
	//log.Printf("%+v\n", resp)
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
	print(search2)
	fmt.Println("Search results are stored in 'print.json' file")
}
func handleSearch(client *twitter.Client, searchCmd *flag.FlagSet, text *string) {
	searchCmd.Parse(os.Args[2:])
	if *text == "" {
		fmt.Print("topic is required to search")
		searchCmd.PrintDefaults()
		os.Exit(1)
	} else {
		search(client, *text)
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

func userTimeline(client *twitter.Client, userTimelineCmd *flag.FlagSet, userTimelineScreenName *string) {
	userTimelineCmd.Parse(os.Args[2:])
	if *userTimelineScreenName == "" {
		fmt.Print("Screen Name is required to see the timeline od the user")
		userTimelineCmd.PrintDefaults()
		os.Exit(1)
	} else {
		timeline, _, err := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
			ScreenName: *userTimelineScreenName,
			//TrimUser:        flag.Bool("faltu", true, "faltu"),
			IncludeRetweets: flag.Bool("faaltu1", false, "faaltu1"),
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
		print(timeline1)
		fmt.Println("Results are stored in 'print.json' file")
	}
}
