package dto

import "time"

type AuthResult struct {
	Success    bool
	Token      string
	Message    string
	Expiration time.Time
}
