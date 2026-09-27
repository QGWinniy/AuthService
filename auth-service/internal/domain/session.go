package domain

import "time"

type Session struct {
	SessionID string
	Jwt string
	UserID    uint64
	ExpiresAt time.Time
}