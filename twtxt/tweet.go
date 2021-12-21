package twtxt

import "time"

type Tweet struct {
	Timestamp time.Time
	Message   string
}
