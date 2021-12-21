package twtxt

type Config struct {
	nick              string
	twtfile           string
	twturl            string
	check_following   bool
	disclose_identity bool
	timeout           float64
}
