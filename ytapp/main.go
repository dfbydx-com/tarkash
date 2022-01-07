package main

import (
	"fmt"

	"github.com/Dilshad-create/ytapp/youtube"
)

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "I won't print hello world this time")
// }

// func setupRoutes() {
// 	http.HandleFunc("/", homePage)
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

func main() {
	fmt.Println("Youtube API tests")
	item, err := youtube.GetSubscribers("UClQf58OQvm6OkGbQARqqhMg")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", item)

	//setupRoutes()
}
