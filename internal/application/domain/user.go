package domain

import "github.com/google/uuid"

type UserID string

type User struct {
	ID             UserID
	UserAgent      string
	IPAddress      string
	IsDeauthorised bool
}

func NewUser(userAgent string, ipAddress string) *User {
	return &User{
		ID:             UserID(uuid.New().String()),
		UserAgent:      userAgent,
		IPAddress:      ipAddress,
		IsDeauthorised: false,
	}
}
