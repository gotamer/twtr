package twtxt

import (
	"net/url"
	"os"
)

// Config is a collection of configuration options that control how to post new
// tweets, as well as retrieving the tweets from followed users.
type Config struct {
	NickName string
	TwtFile  *os.File
	TwtUrl   *url.URL
	Timeout  float64
}
