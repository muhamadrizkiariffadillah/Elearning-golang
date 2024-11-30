package users

import "time"

type Users struct {
	Id           int
	FullName     string
	Username     string
	Email        string
	HashPassword string
	Role         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
