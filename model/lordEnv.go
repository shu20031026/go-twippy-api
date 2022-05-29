package model

import (
	"os"
)

func LoadEnv() (string, string, string, string) {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	fmt.Printf(".envファイルの読み込みに失敗しました: %v", err)
	// }
	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")

	return consumerKey, consumerSecret, accessToken, accessTokenSecret
}
