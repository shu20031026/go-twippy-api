package model

import (
	"fmt"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/joho/godotenv"
)

var consumerKey string
var consumerSecret string
var accessToken string
var accessTokenSecret string

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf(".envファイルの読み込みに失敗しました: %v", err)
	}
	consumerKey = os.Getenv("CONSUMER_KEY")
	consumerSecret = os.Getenv("CONSUMER_SECRET")
	accessToken = os.Getenv("ACCESS_TOKEN")
	accessTokenSecret = os.Getenv("ACCESS_TOKEN_SECRET")
}

func FetchTweet(name string) []anaconda.Tweet {
	LoadEnv()

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

// // 型定義
// type (
// 	tweetDataType struct {
// 		Name       string   `json:"name"`
// 		ScreenName string   `json:"screenName"`
// 		Icon       string   `json:"icon"`
// 		Tweets     []string `json:"tweets"`
// 	}
// )
