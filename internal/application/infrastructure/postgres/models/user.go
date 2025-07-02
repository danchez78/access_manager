package models

import "access_manager/internal/application/domain"

type User struct {
	ID             string
	UserAgent      string
	IPAddress      string
	IsDeauthorised bool
}

func UserFromDomain(user *domain.User) *User {
	return &User{
		ID:             string(user.ID),
		UserAgent:      user.UserAgent,
		IPAddress:      user.IPAddress,
		IsDeauthorised: user.IsDeauthorised,
	}
}

func (u *User) ToDomain() *domain.User {
	return &domain.User{
		ID:             domain.UserID(u.ID),
		UserAgent:      u.UserAgent,
		IPAddress:      u.IPAddress,
		IsDeauthorised: u.IsDeauthorised,
	}
}
