package domain

import "time"

type Video struct {
	ID         string // internal identification
	ResourceID string // id from client
	FilePath   string // bucket`s location
	CreatedAt  time.Time
}
