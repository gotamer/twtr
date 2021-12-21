package twtxt

import (
	"os"
)

// Config is a collection of configuration options that control how to post new
// tweets, as well as retrieving the tweets from followed users.
type Config struct {
	Account Account
	TwtFile *os.File
	Timeout float64
}
