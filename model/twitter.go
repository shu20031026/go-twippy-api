package model

import (
	"fmt"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

func FetchTweet(name string, fetchCount string) []anaconda.Tweet {
	//初期化
	consumerKey, consumerSecret, accessToken, accessTokenSecret := LoadEnv()
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	twitter := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	// クエリパラメータの設定
	timelineQuery := url.Values{}
	timelineQuery.Set("screen_name", name)
	timelineQuery.Set("count", fetchCount)
	timelineQuery.Set("include_rts", "false")
	timelineQuery.Set("exclude_replies", "false")
	fetchData, err := twitter.GetUserTimeline(timelineQuery)

	if err != nil {
		fmt.Printf("Error to getUserTimeline. err:%v\n", err)
		os.Exit(1)
	}

	return fetchData
}
