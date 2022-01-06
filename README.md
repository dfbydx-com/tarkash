# HACKS-TOOLS
social media tools and other hacks
1. twapp
2. ytapp
3. linkapp

We are still working on ytapp, linkapp. Here is the readme for twapp.
# Introduction
The twitter command-line tool developed in golang lets us do some awesome things:
   * search the topics, hashtags on twitter
   * view timline of a user by their screen name
   * save the topics
   * save the users
   * search for all the saved topics
   * view timeline for all  the saved users

# Getting the code
The code is hosted at https://github.com/mypass-id/hacks-tools

Check out the latest development version anonymously with:
```
$ git clone https://github.com/mypass-id/hacks-tools
$ cd twapp
```

# Understanding files
## twapp.exe
We have `twapp.exe` file which we can make by running
```
$ go build
```
## print.json
Everytime we call the apis and get the results from twitter, we save the usefull parts of result in `print.json` file. It's better way in case we want other formats.
## config.json
We have filled the file with dummy data. You can edit the file with your credentials to call twitter api. We will see how to get the credentials later.
Dummy data:
```
{
    "AccessToken" :    "147334...",
    "AccessTokenSecret": "1nMXfjOt...",
	"ConsumerKey":      "n2d10jp...",
	"ConsumerSecret":   "wBztEnFQwC..."
}
```

## topics.json
In this file, we have added all the topics we may need to search in future. We can add, delete, get it's data in cli using commands we have build.
Dummy data:
```
[
    {
        "Text": "technology",
        "Added_at": "2021-12-31 01:28:24.3979122 +0530 IST m=+0.567789001"
    },
    {
        "Text": "#covid",
        "Added_at": ""
    }
]
```

## users.json
In this file, we have added all the users we want to focus on. We can add, delete, get it's data in cli using commands we have build.
Dummy data:
```
[
    {
        "Screen_Name": "elonmusk",
        "Added_at": "2021-12-31 01:30:38.014643 +0530 IST m=+0.585155001"
    },
    {
        "Screen_Name": "jack",
        "Added_at": "2021-12-31 01:30:53.4396434 +0530 IST m=+0.589103201"
    }
]
```
# Authentication
To authenticate, visit the Twitter developer page and create a new application:
https://dev.twitter.com/apps/new

***
**Get your a keys and secrets.
Edit the `config.json` file with your credentials like this:**
```
{
    "AccessToken" :    "14733451...",
    "AccessTokenSecret": "1nMXfjOt...",
	"ConsumerKey":      "n2d10jpK...",
	"ConsumerSecret":   "wBztEnF..."
} 
```
***

By design, the Twitter Client accepts any http.Client so user auth (OAuth1) or application auth (OAuth2) requests can be made by using the appropriate authenticated client. Use the https://github.com/dghubble/oauth1 and https://github.com/golang/oauth2 packages to obtain an http.Client which transparently authorizes requests.

Finally, we can use the OAuth1 authenticator to connect to Twitter. In code it all goes like this:

```
import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

config := oauth1.NewConfig("consumerKey", "consumerSecret")
token := oauth1.NewToken("accessToken", "accessSecret")
// http.Client will automatically authorize Requests
httpClient := config.Client(oauth1.NoContext, token)

// twitter client
client := twitter.NewClient(httpClient)
```
Similarly, we can use the OAuth2 authenticator and token to connect to Twitter. In code it goes like this:
```
import (
	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2"
)

config := &oauth2.Config{}
token := &oauth2.Token{AccessToken: accessToken}
// http.Client will automatically authorize Requests
httpClient := config.Client(oauth2.NoContext, token)

// twitter client
client := twitter.NewClient(httpClient)
```

See https://developer.twitter.com/en/docs for more detail.

# Commands

## search for topic
* `twapp search --all`

   search for all the already added topics in `topics.json`
   
* `twapp search -topic "bollywood"`
   
   search for specific topic
* `twapp search -topic "#covid"`
## view user-timeline
* `twapp usertimeline --all` 

   view timline for the already added ussers in `users.json`
* `twapp usertimeline -screenName "elonmusk"`
   
   timeline for specific user

## topics.json related commands
* `twapp topics get -all`

   view all topics added in `topics.json` till now
* `twapp topics get -topic "bollywood"`

   view specific topic, return *topic not available* if topic is not available in `topics.json`
* `twapp topics add -topic "bollywood"`

   add "bollywood" in `topics.json`, updates the added_at field if already added.
* `twapp topics remove -topic "#covid"`

   remove "#covid" from `topics.json`, return *topic not available* if topic is not available in `topics.json`
## users.json related commands
* `twapp users get -all`

   view all added users in `users.json` till now
* `twapp users get -user "jack"`

   view specific user, return *user not available* if topic is not available in `users.json`
* `twapp users add -user "elonmusk"`

   add "elonmusk" in `users.json`, updates the added_at field if already added.
* `twapp users remove -user "jack"`

   remove "jack" from `users.json`, return *user not available* if user is not available in `users.json`
   