package controller

import (
	"main/model"
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func GetEndPoint(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, twippy api")
}

// 型定義
type MoldDataType struct {
	Name       string   `json:"name"`
	ScreenName string   `json:"screenName"`
	Icon       string   `json:"icon"`
	Tweets     []string `json:"tweets"`
}

func GetTweets(c echo.Context) error {
	name := c.QueryParam("name")

	FETCH_COUNT := "100"

	tweets := model.FetchTweet(name, FETCH_COUNT)

	var moldData MoldDataType
	moldData.Icon = tweets[0].User.ProfileImageUrlHttps
	moldData.Name = tweets[0].User.Name
	moldData.ScreenName = tweets[0].User.ScreenName
	var TweetsSlice = moldData.Tweets

	for _, tweet := range tweets {
		TweetsSlice = append(TweetsSlice, tweet.FullText)
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(TweetsSlice), func(i, j int) { TweetsSlice[i], TweetsSlice[j] = TweetsSlice[j], TweetsSlice[i] })

	ReturnTweets := TweetsSlice[0:5]

	moldData.Tweets = ReturnTweets

	return c.JSON(http.StatusOK, moldData)
}
