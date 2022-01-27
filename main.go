package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	tweetDataType struct {
		Name       string   `json:"name"`
		ScreenName string   `json:"screenName"`
		Icon       string   `json:"icon"`
		Tweets     []string `json:"tweets"`
	}
)

var consumerKey = ""
var consumerSecret = ""
var accessToken = ""
var accessTokenSecret = ""

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	twitter := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	timelineQuery := url.Values{}

	timelineQuery.Set("screen_name", "fukke0906")
	timelineQuery.Set("count", "20")
	timelineQuery.Set("include_rts", "false")
	fmt.Println(timelineQuery)
	tweets, err := twitter.GetUserTimeline(timelineQuery)
	if err != nil {
		fmt.Printf("Error to getUserTimeline. err:%v\n", err)
		os.Exit(1)
	}

	fmt.Println(tweets)
	e.GET("/", hello)
	e.GET("/test", test)
	e.GET("/tweets", moldData)

	e.Logger.Fatal(e.Start(":80"))
}

// ハンドラーを定義
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, twippy-api")
}

func test(c echo.Context) error {
	return c.String(http.StatusOK, "test")
}

func moldData(c echo.Context) error {
	return c.JSON(http.StatusOK, tweetDataType{})
}
