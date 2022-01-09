package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Dilshad-create/ytapp/youtube"
)

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "I don't want to write 'hello world' this time")
// }

// func setupRoutes() {
// 	http.HandleFunc("/", homePage)
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

func main() {
	fmt.Println("Youtube API tests")

	searchCmd := flag.NewFlagSet("search", flag.ExitOnError)
	searchQuery := searchCmd.String("query", "", "What to search?")
	searchMaxResults := searchCmd.String("maxResults", "", "OPTIONAL PARAMETER, MaxResults per page?")

	subscriberCmd := flag.NewFlagSet("channel", flag.ExitOnError)
	subscriberId := subscriberCmd.String("id", "", "What's the id of the channel?")

	if len(os.Args) < 2 {
		fmt.Println("expected 'search' or 'channel' subcommands")
		os.Exit(1)
	}

	//look at the 2nd argument's value
	switch os.Args[1] {
	case "search":
		handleSearch(searchCmd, searchQuery, searchMaxResults)
	case "channel":
		handleSubscriber(subscriberCmd, subscriberId)
	default:
		fmt.Println("we have 'search' and 'channel' subcommands only")

	}

	//setupRoutes()
}

func handleSearch(searchCmd *flag.FlagSet, query *string, maxResults *string) {
	searchCmd.Parse(os.Args[2:])
	if *query == "" {
		fmt.Println("query is required to search")
		searchCmd.PrintDefaults()
		os.Exit(1)
	}
	if *maxResults == "" {
		*maxResults = "5"
	}
	resp, err := youtube.Search(*query, *maxResults) //attributes: query,maxResults
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf("%+v\n", resp)
	youtube.PrintSearch(resp)
}

func handleSubscriber(subscriberCmd *flag.FlagSet, id *string) {
	subscriberCmd.Parse(os.Args[2:])
	if *id == "" {
		fmt.Println("channel id is required")
		subscriberCmd.PrintDefaults()
		os.Exit(1)
	}

	// "UCt4-7kmQaPEZzPLil4RNRCw",youtube channel id of iit bhu fmc.
	item, err := youtube.GetSubscribers(*id) //attributes:channel id
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf("%+v\n", item)
	youtube.PrintGetSubscribers(item)
}
