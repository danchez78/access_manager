package routes

import (
	"access_manager/internal/application/usecases"
	"access_manager/internal/common/server"
)

func Make(
	srv *server.Server,
	uc *usecases.UseCases,
) {
	makeAccessRoutes(srv, uc)
}
