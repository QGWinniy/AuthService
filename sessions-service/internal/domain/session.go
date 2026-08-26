package domain

import "time"

type Session struct {
	SessionId string
	UserId uint64
	Data string
	ExpiresAt time.Time
}