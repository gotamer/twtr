package twtxt

import "net/url"

type Account struct {
	NickName string
	TwtxtUrl *url.URL
}
