package model

import (
	"fmt"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

func FetchTweet(name string) []anaconda.Tweet {
	consumerKey, consumerSecret, accessToken, accessTokenSecret := LoadEnv()

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	twitter := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	timelineQuery := url.Values{}

	timelineQuery.Set("screen_name", name)
	timelineQuery.Set("count", "100")
	timelineQuery.Set("include_rts", "false")
	timelineQuery.Set("exclude_replies", "false")
	tweets, err := twitter.GetUserTimeline(timelineQuery)

	if err != nil {
		fmt.Printf("Error to getUserTimeline. err:%v\n", err)
		os.Exit(1)
	}

	return tweets
}
