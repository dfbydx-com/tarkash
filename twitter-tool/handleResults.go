package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func print(tweets []tweet) {
	tweetBytes, err := json.Marshal(tweets)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("./print.json", tweetBytes, 0644)
	if err != nil {
		panic(err)
	}
}

type topic struct {
	Text string
	//Added_at time.Time
	Added_at string
}
type user struct {
	Screen_Name string
	Added_at    string
	//other twitter.User datas
}

func getTopics() (topics []topic) {
	fileBytes, err := ioutil.ReadFile("topics.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(fileBytes, &topics)
	if err != nil {
		fmt.Println("Maybe 'topics.json' file was missing, or it is empty")
	}
	return topics
}
func saveTopics(topics []topic) {
	topicBytes, err := json.Marshal(topics)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("./topics.json", topicBytes, 0644)
	if err != nil {
		panic(err)
	}
}
func getUsers() (users []user) {
	fileBytes, err := ioutil.ReadFile("users.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(fileBytes, &users)
	if err != nil {
		fmt.Println("Maybe 'users.json' file was missing, or it is empty")
	}
	return users
}
func saveUsers(users []user) {
	userBytes, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("./users.json", userBytes, 0644)
	if err != nil {
		panic(err)
	}
}

func handleGetTopic(topicsGetCmd *flag.FlagSet, topicsGetAll *bool, topicsGetTopic *string) {
	topicsGetCmd.Parse(os.Args[3:])
	if !*topicsGetAll && *topicsGetTopic == "" {
		fmt.Print("topic is required or specify --all to get all the videos")
		topicsGetCmd.PrintDefaults()
		os.Exit(1)
	}
	if *topicsGetAll {
		topics := getTopics()
		fmt.Printf(" TOPIC \t\t\t ADDED_AT \n")
		for _, topic := range topics {
			fmt.Printf("%v \t\t %v\n", topic.Text, topic.Added_at)
		}
		return
	}
	if *topicsGetTopic != "" {
		topics := getTopics()
		found := 0
		for _, topic := range topics {
			if *topicsGetTopic == topic.Text {
				fmt.Printf(" TOPIC \t\t\t ADDED_AT \n")
				fmt.Printf("%v \t\t %v\n", topic.Text, topic.Added_at)
				found = 1
			}
		}
		if found == 0 {
			fmt.Printf("'%v' doesn't exist in topics\n", *topicsGetTopic)
		}
	}
}
func handleAddTopic(topicsAddCmd *flag.FlagSet, topicsAddTopic *string) {
	topicsAddCmd.Parse(os.Args[3:])
	if *topicsAddTopic == "" {
		fmt.Print("-topic is required to add it")
		topicsAddCmd.PrintDefaults()
		os.Exit(1)
	}
	t := topic{
		Text:     *topicsAddTopic,
		Added_at: time.Now().String(),
	}
	topics := getTopics()
	for i, topic := range topics {
		if *topicsAddTopic == topic.Text {
			fmt.Println("requested topic is already added, updated 'Added_at' field")
			topics[i] = t
			return
		}
	}
	topics = append(topics, t)
	saveTopics(topics)
}
func handleRemoveTopic(topicsRemoveCmd *flag.FlagSet, topicsRemoveTopic *string) {
	topicsRemoveCmd.Parse(os.Args[3:])
	if *topicsRemoveTopic == "" {
		fmt.Print("-topic is required to remove it")
		topicsRemoveCmd.PrintDefaults()
		os.Exit(1)
	}
	topics := getTopics()
	for i, topic := range topics {
		if *topicsRemoveTopic == topic.Text {
			fmt.Println("OK, requested topic is removed")
			t := append(topics[:i], topics[i+1:]...)
			saveTopics(t)
			return
		}
	}
	fmt.Println("Requested topic is NOT FOUND")
}
func handleGetUser(usersGetCmd *flag.FlagSet, usersGetAll *bool, usersGetUser *string) {
	usersGetCmd.Parse(os.Args[3:])
	if !*usersGetAll && *usersGetUser == "" {
		fmt.Print("topic is required or specify --all to get all the videos")
		usersGetCmd.PrintDefaults()
		os.Exit(1)
	}
	if *usersGetAll {
		users := getUsers()
		fmt.Printf(" SCREEN_NAME \t\t\t ADDED_AT \n")
		for _, user := range users {
			fmt.Printf("%v \t\t %v\n", user.Screen_Name, user.Added_at)
		}
		return
	}
	if *usersGetUser != "" {
		users := getUsers()
		found := 0
		for _, user := range users {
			if *usersGetUser == user.Screen_Name {
				fmt.Printf(" SCREEN_NAME \t\t\t ADDED_AT \n")
				fmt.Printf("%v \t\t %v\n", user.Screen_Name, user.Added_at)
				found = 1
			}
		}
		if found == 0 {
			fmt.Printf("'%v' doesn't exist in users\n", *usersGetUser)
		}
	}
}
func handleAddUser(usersAddCmd *flag.FlagSet, usersAddUser *string) {
	usersAddCmd.Parse(os.Args[3:])
	if *usersAddUser == "" {
		fmt.Print("-topic is required to add it")
		usersAddCmd.PrintDefaults()
		os.Exit(1)
	}
	t := user{
		Screen_Name: *usersAddUser,
		Added_at:    time.Now().String(),
	}
	users := getUsers()
	//********************************************************
	// user is already defined as struct,,,fix this if it is an issue
	//********************************************************
	for i, user := range users {
		if *usersAddUser == user.Screen_Name {
			fmt.Println("requested user is already added, updated 'Added_at' field")
			users[i] = t
			return
		}
	}
	users = append(users, t)
	saveUsers(users)
}
func handleRemoveUser(usersRemoveCmd *flag.FlagSet, usersRemoveUser *string) {
	usersRemoveCmd.Parse(os.Args[3:])
	if *usersRemoveUser == "" {
		fmt.Print("-topic is required to remove it")
		usersRemoveCmd.PrintDefaults()
		os.Exit(1)
	}
	users := getUsers()
	for i, user := range users {
		if *usersRemoveUser == user.Screen_Name {
			fmt.Println("OK, requested topic is removed")
			t := append(users[:i], users[i+1:]...)
			saveUsers(t)
			return
		}
	}
	fmt.Println("Requested user is NOT FOUND")
}
