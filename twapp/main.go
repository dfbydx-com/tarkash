package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	ConsumerKey       string //api key
	ConsumerSecret    string //api secret
	AccessToken       string
	AccessTokenSecret string
}

func main() {
	creds := Configuration{}
	err := gonfig.GetConf("./config.json", &creds)
	if err != nil {
		log.Printf("Error getting credentials from config")
		return
	}
	client, err := TwitterCredentialsCheck(&creds)
	if err != nil {
		log.Println("Error getting Twitter Client")
		log.Println(err)
		return
	}

	//-----------------------
	// api search --all
	// api search -topic "bollywood"
	searchCmd := flag.NewFlagSet("search", flag.ExitOnError)
	searchAll := searchCmd.Bool("all", false, "search for all added topics?")
	searchString := searchCmd.String("topic", "", "What to search?")
	//-----------------------

	// api topics get -all (search for all added topics in topics.json)
	// api topics get -topic "bollywood"
	// api topics add -topic "bollyhwood"
	// api topics remove -topic "bollywood"
	topicsGetCmd := flag.NewFlagSet("get", flag.ExitOnError)
	topicsAddCmd := flag.NewFlagSet("add", flag.ExitOnError)
	topicsRemoveCmd := flag.NewFlagSet("remove", flag.ExitOnError)
	topicsGetAll := topicsGetCmd.Bool("all", false, "Get All Topics")
	topicsGetTopic := topicsGetCmd.String("topic", "", "Which topic you want to see?")
	topicsAddTopic := topicsAddCmd.String("topic", "", "Which topic to add?")
	topicsRemoveTopic := topicsRemoveCmd.String("topic", "", "Which topic to remove?")
	//------------------------

	// api usertimeline --all  (timeline for added users in users.json)
	// api usertimeline -screenName "elonmusk"
	userTimelineCmd := flag.NewFlagSet("usertimeline", flag.ExitOnError)
	userTimelineAll := userTimelineCmd.Bool("all", false, "user timeline for all the added users?")
	userTimelineScreenName := userTimelineCmd.String("screenName", "", "what is the screen name of the user?")
	//------------------------

	// api users get -all
	// api users get -user "elonmusk"
	// api users add -user "elonmusk"
	// api users remove -user "elonmusk"
	usersGetCmd := topicsGetCmd
	usersAddCmd := topicsAddCmd
	usersRemoveCmd := topicsRemoveCmd
	usersGetAll := topicsGetAll
	usersGetUser := usersGetCmd.String("user", "", "Which user you want to see?")
	usersAddUser := usersAddCmd.String("user", "", "Which user to add?")
	usersRemoveUser := usersRemoveCmd.String("user", "", "Which user to remove?")
	//------------------------

	// api trends
	flag.NewFlagSet("trends", flag.ExitOnError)
	//------------------------

	// api followers -id 12 (id of the user)
	followersCmd := flag.NewFlagSet("followers", flag.ExitOnError)
	followersId := followersCmd.Int64("id", 0, "What's the user id?")
	//-------------------------

	if len(os.Args) < 2 {
		fmt.Println("expected 'search', 'usertimeline', 'topics', 'users', 'trends' or 'followers' subcommands")
		os.Exit(1)
	}

	//look at the 2nd argument's value
	switch os.Args[1] {
	case "search":
		handleSearch(client, searchCmd, searchAll, searchString)
	case "usertimeline":
		handleUserTimeline(client, userTimelineCmd, userTimelineAll, userTimelineScreenName)
	case "topics":
		if len(os.Args) < 3 {
			fmt.Println("expected 'get', 'add' or 'remove' subcommands after topics")
			os.Exit(1)
		} else if os.Args[2] == "get" {
			handleGetTopic(topicsGetCmd, topicsGetAll, topicsGetTopic)
		} else if os.Args[2] == "add" {
			handleAddTopic(topicsAddCmd, topicsAddTopic)
		} else if os.Args[2] == "remove" {
			handleRemoveTopic(topicsRemoveCmd, topicsRemoveTopic)
		} else {
			fmt.Println("expected 'get', 'add' or 'remove' subcommands after topics")
			os.Exit(1)
		}
	case "users":
		if len(os.Args) < 3 {
			fmt.Println("expected 'get', 'add' or 'remove' subcommands after users")
			os.Exit(1)
		} else if os.Args[2] == "get" {
			handleGetUser(usersGetCmd, usersGetAll, usersGetUser)
		} else if os.Args[2] == "add" {
			handleAddUser(usersAddCmd, usersAddUser)
		} else if os.Args[2] == "remove" {
			handleRemoveUser(usersRemoveCmd, usersRemoveUser)
		} else {
			fmt.Println("expected 'get', 'add' or 'remove' subcommands after users")
			os.Exit(1)
		}
	case "trends":
		trends(client)
	case "followers":
		followers(client, followersCmd, followersId)
	default:
		fmt.Println("expected 'search', 'usertimeline', 'topics', 'users', 'trends' or 'followers' subcommands")
	}

}
