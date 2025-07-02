package usecases

import (
	"access_manager/internal/application/domain"
	"context"
)

type CreateUserHandler struct {
	userRepo domain.UserRepository
}

func NewCreateUserHandler(userRepo domain.UserRepository) *CreateUserHandler {
	if userRepo == nil {
		panic("user repository is nil")
	}
	return &CreateUserHandler{userRepo: userRepo}
}

func (h *CreateUserHandler) Execute(ctx context.Context, userAgent string, ipAddress string) (domain.UserID, error) {
	user := domain.NewUser(userAgent, ipAddress)
	if err := h.userRepo.Save(ctx, user); err != nil {
		return "", err
	}
	return user.ID, nil
}
