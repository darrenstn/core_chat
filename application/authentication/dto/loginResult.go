package dto

import "time"

type LoginResult struct {
	Success    bool
	Token      string
	Message    string
	Expiration time.Time
}
