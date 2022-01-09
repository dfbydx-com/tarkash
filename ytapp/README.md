# Introduction
In ytapp, we have implemented Search(), GetSubscribers() using youtbe apis.
   * search the query, get related videos, channels etc.
   * get channel data like view count, subscriber count. 

# Getting the code
The code is hosted at https://github.com/mypass-id/hacks-tools

Check out the latest development version anonymously with:
```
$ git clone https://github.com/mypass-id/hacks-tools
$ cd ytapp
```

# Understanding files
## ytapp.exe
We need `ytapp.exe` file to run the commands.which we can make by running
```
$ go build
```
## print.json
Everytime we call the apis and get the results from youtube, we save the usefull parts of result in `print.json` file. It's better way in case we want other formats.
## config.json
We have filled the file with dummy data. You can edit the file with your credentials to call youtube api. We will see how to get the credentials later.
Dummy data:
```
{
    "AccessToken" :    "147334...",
    "AccessTokenSecret": "1nMXfjOt...",
	"ConsumerKey":      "n2d10jp...",
	"ConsumerSecret":   "wBztEnFQwC..."
}
```
# Authentication
To authenticate, visit the https://developers.google.com/youtube/v3/docs page and follow instructions:

***
**Get your a keys and secrets. At this stage, we need only Api Key.
Edit the `config.json` file with your credentials like this:**
```
{
    "Key" :    "AIz..."
}
```


# Run
We have implemented Search(), GetSubscribers() funcs in youtube package.

## Search()
It takes query and maxResults attributes. 

* The query parameter specifies the query term to search for.
* The maxResults parameter specifies the maximum number of items that should be returned in the result set. Acceptable values are 0 to 50, inclusive. The default value is 5.

**Result for Search("new year","2"):**
```
{
    "kind": "youtube#searchListResponse",
    "nextPageToken": "CAIQAA",
    "pageInfo": {
        "totalResults": 1000000,
        "resultsPerPage": 2
    },
    "items": [
        {
            "kind": "youtube#searchResult",
            "id": {
                "kind": "youtube#video",
                "videoId": "ZKxC_nS0870"
            },
            "snippet": {
                "publishedAt": "2022-01-03T03:00:09Z",
                "channelId": "UCjIwesauyswaliQSQE6syig",
                "title": "Happy New Year Special New Comedy Video 2022 Amazing Funny Video 2021 Episode 136 By Funny Day",
                "description": "In this video You are watching,Must Watch New Funny Video 2021_Top New Comedy Video 2021_Try To Not Laugh ...",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/ZKxC_nS0870/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/ZKxC_nS0870/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/ZKxC_nS0870/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "Funny Day"
            }
        },
        {
            "kind": "youtube#searchResult",
            "id": {
                "kind": "youtube#video",
                "videoId": "lwLSb1cq4gM"
            },
            "snippet": {
                "publishedAt": "2021-12-31T19:35:14Z",
                "channelId": "UCnoqvTW4YZExfDeq7_Wmd-w",
                "title": "Pellam Vaddu Party Muddu | Extra Jabardasth|ETV New Year Special Event| Full Episode |31st Dec21|ETV",
                "description": "rgv #extrajabardasth #pellamvaddupartymuddu #etvnewyearevent2022 #etvnewyearevents #etvevents #etvshows #indraja ...",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/lwLSb1cq4gM/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/lwLSb1cq4gM/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/lwLSb1cq4gM/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "ETV Jabardasth"
            }
        }
    ]
}
```

## GetSubscribers()
The id parameter specifies a YouTube channel ID for the resource(s) that are being retrieved.
**Result for GetSubscribers("UCt4-7kmQaPEZzPLil4RNRCw"):**
```
{
    "kind": "youtube#channel",
    "id": "UCt4-7kmQaPEZzPLil4RNRCw",
    "statistics": {
        "viewCount": "1212732",
        "subscriberCount": "8720"
    }
}
```
