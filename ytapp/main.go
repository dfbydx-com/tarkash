package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Dilshad-create/ytapp/youtube"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "I don't want to write 'hello world' this time")
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	fmt.Println("Youtube API tests")
	// item, err := youtube.GetSubscribers("UClQf58OQvm6OkGbQARqqhMg") //attributes:channel id
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// //fmt.Printf("%+v\n", item)
	// youtube.PrintGetSubscribers(item)

	resp, err := youtube.Search("new year", "5") //attributes: query,maxResults
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf("%+v\n", resp)
	youtube.PrintSearch(resp)

	//setupRoutes()
}
