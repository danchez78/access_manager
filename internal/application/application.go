package application

import (
	"context"

	"access_manager/config"
	"access_manager/internal/application/domain"
	"access_manager/internal/application/infrastructure/api/routes"
	"access_manager/internal/application/infrastructure/postgres"
	"access_manager/internal/application/usecases"
	"access_manager/internal/common/alert_sender"
	"access_manager/internal/common/postgres_client"
	"access_manager/internal/common/server"
	"access_manager/internal/common/token_generator"
)

func Init(
	ctx context.Context,
	srv *server.Server,
	cfg config.Config,
) error {
	psql_client, err := postgres_client.NewClient(ctx, cfg.Postgres)
	if err != nil {
		return err
	}

	tg := token_generator.NewTokenGenerator(cfg.TokenGenerator)
	tm := domain.NewTokenManager(tg)

	alert_s := alert_sender.NewAlertSender(cfg.AlertSender)
	as := domain.NewAlertSender(alert_s)

	userRepo := postgres.NewUserRepository(psql_client)
	tokenRepo := postgres.NewTokenRepository(psql_client)

	uc := usecases.NewUseCases(userRepo, tokenRepo, tm, as)
	routes.Make(srv, uc)

	return nil
}
