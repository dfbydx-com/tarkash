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

type SearchResp struct {
	Kind          string       `json:"kind"`
	NextPageToken string       `json:"nextPageToken"`
	PageInfo      PageInfo     `json:"pageInfo"`
	Items         []SearchItem `json:"items"`
}
type PageInfo struct {
	TotalResults   int64 `json:"totalResults"`
	ResultsPerPage int64 `json:"resultsPerPage"`
}
type SearchItem struct {
	Kind    string        `json:"kind"`
	Id      SearchId      `json:"id"`
	Snippet SearchSnippet `json:"snippet"`
}
type SearchId struct {
	Kind    string `json:"kind"`
	VideoId string `json:"videoId"`
}
type SearchSnippet struct {
	PulishedAt   string     `json:"publishedAt"`
	ChannelId    string     `json:"channelId"`
	Title        string     `json:"title"`
	Desc         string     `json:"description"`
	Thumbnails   Thumbnails `json:"thumbnails"`
	ChannelTitle string     `json:"channelTitle"`
}
type Thumbnails struct {
	Default Thumbnail `json:"default"`
	Medium  Thumbnail `json:"medium"`
	High    Thumbnail `json:"high"`
}
type Thumbnail struct {
	Url    string `json:"url"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
}

func Search(query string, maxResults string) (SearchResp, error) {
	response := SearchResp{}

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://youtube.googleapis.com/youtube/v3/search", nil)
	if err != nil {
		fmt.Println(err)
		return response, err
	}

	creds := YoutubeCreds{}
	err = gonfig.GetConf("./youtube/config.json", &creds)
	if err != nil {
		fmt.Println(err)
		return response, err
	}

	q := req.URL.Query()
	q.Add("part", "snippet")
	q.Add("q", query) // new year should be new%20year
	q.Add("key", creds.Key)
	q.Add("maxResults", maxResults)

	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return response, err
	}
	defer resp.Body.Close()

	fmt.Println("Response Status: ", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return response, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return response, err
	}

	return response, nil
}

func PrintSearch(response SearchResp) {
	bytes, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("./print.json", bytes, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Result is saved in './print.json' ")
}
func PrintGetSubscribers(item Item) {
	bytes, err := json.Marshal(item)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("./print.json", bytes, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Result is saved in './print.json' ")
}
