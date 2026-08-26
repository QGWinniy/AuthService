package domain

import "time"

type Session struct {
	SessionID string
	UserID    uint64
	ExpiresAt time.Time
}