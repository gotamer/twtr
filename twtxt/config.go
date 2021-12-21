package twtxt

// Config is a collection of configuration options that control how to post new
// tweets, as well as retrieving the tweets from followed users.
type Config struct {
	nick              string
	twtfile           string
	twturl            string
	check_following   bool
	disclose_identity bool
	timeout           float64
}
