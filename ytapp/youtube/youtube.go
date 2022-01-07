package youtube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/tkanos/gonfig"
)

type Response struct {
	Kind  string `json:"kind"`
	Items []Item `json:"items"`
}
type Item struct {
	Kind  string `json:"kind"`
	Id    string `json:"id"`
	Stats Stats  `json:"statistics"`
}
type Stats struct {
	Views       string `json:"viewCount"`
	Subscribers string `json:"subscriberCount"`
}

type YoutubeCreds struct {
	Key               string //api key
	AccessToken       string
	AccessTokenSecret string
}

func GetSubscribers(id string) (Item, error) {
	var response Response

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://youtube.googleapis.com/youtube/v3/channels", nil)
	if err != nil {
		fmt.Println(err)
		return Item{}, err
	}

	creds := YoutubeCreds{}
	err = gonfig.GetConf("./youtube/config.json", &creds)
	if err != nil {
		fmt.Println(err)
		return Item{}, err
	}

	q := req.URL.Query()
	q.Add("key", creds.Key)
	q.Add("id", id)
	q.Add("part", "statistics")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return Item{}, err
	}
	defer resp.Body.Close()

	fmt.Println("Response Status: ", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return Item{}, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return Item{}, err
	}

	return response.Items[0], nil
}
