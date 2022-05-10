package main

import "main/router"

func main() {
	router.Init()
}

// 初期化
// func loadEnv() {
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		fmt.Printf(".envファイルの読み込みに失敗しました: %v", err)
// 	}
// 	consumerKey = os.Getenv("CONSUMER_KEY")
// 	consumerSecret = os.Getenv("CONSUMER_SECRET")
// 	accessToken = os.Getenv("ACCESS_TOKEN")
// 	accessTokenSecret = os.Getenv("ACCESS_TOKEN_SECRET")
// }

// func moldData(c echo.Context) error {
// 	anaconda.SetConsumerKey(consumerKey)
// 	anaconda.SetConsumerSecret(consumerSecret)
// 	twitter := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

// 	timelineQuery := url.Values{}

// 	name := c.QueryParam("name")

// 	timelineQuery.Set("screen_name", name)
// 	timelineQuery.Set("count", "100")
// 	timelineQuery.Set("include_rts", "false")
// 	timelineQuery.Set("exclude_replies", "false")
// 	tweets, err := twitter.GetUserTimeline(timelineQuery)

// 	if err != nil {
// 		fmt.Printf("Error to getUserTimeline. err:%v\n", err)
// 		os.Exit(1)
// 	}

// 	// fmt.Println(tweets[0])

// 	return c.JSON(http.StatusOK, tweets)
// }

// // 型定義
// type (
// 	tweetDataType struct {
// 		Name       string   `json:"name"`
// 		ScreenName string   `json:"screenName"`
// 		Icon       string   `json:"icon"`
// 		Tweets     []string `json:"tweets"`
// 	}
// )
