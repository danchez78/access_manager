package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"access_manager/internal/application/infrastructure/api/controllers"
	"access_manager/internal/application/infrastructure/api/views"
	"access_manager/internal/application/usecases"
	"access_manager/internal/common/server"
)

type accessHandler struct {
	uc *usecases.UseCases
}

// createUser godoc
//
//	@Summary		Create user
//	@Description	Creates user with guid
//	@Tags			access
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	server.Result[views.CreateUserResponse]
//	@Router			/users [post]
func (h *accessHandler) createUser(c echo.Context) error {
	ctx := c.Request().Context()

	userAgent := c.Request().Header.Get("User-Agent")

	userID, err := h.uc.CreateUserHandler.Execute(ctx, userAgent, c.RealIP())
	if err != nil {
		err = fmt.Errorf("failed to create user. Reason: %v", err)
		log.Println(err)

		return server.ReturnError(c, http.StatusInternalServerError, err)
	}

	return server.ReturnResult(c, views.NewCreateUserResponse(userID))
}

// getTokens godoc
//
//	@Summary		Get tokens
//	@Description	Returns access and refresh tokens by user id
//	@Tags			access
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path		string	true	"User ID"
//	@Success		200		{object}	server.Result[views.GetTokensResponse]
//	@Router			/users/{user_id}/tokens [post]
func (h *accessHandler) getTokens(c echo.Context) error {
	ctx := c.Request().Context()

	var req controllers.GetTokensRequest
	if err := c.Bind(&req); err != nil {
		err = fmt.Errorf("failed to decode request. Reason: %v", err)
		log.Println(err)

		return server.ReturnError(c, http.StatusInternalServerError, err)
	}

	accessToken, refreshToken, err := h.uc.GetTokensHandler.Execute(ctx, req.UserID)
	if err != nil {
		err = fmt.Errorf("failed to get tokens. Reason: %v", err)
		log.Println(err)

		return server.ReturnError(c, http.StatusInternalServerError, err)
	}

	return server.ReturnResult(c, views.NewGetTokensResponse(accessToken, refreshToken))
}

// refreshTokens godoc
//
//	@Summary		Refresh tokens
//	@Description	Returns access and refresh tokens by previous access and refresh tokens
//	@Tags			access
//	@Accept			json
//	@Produce		json
//	@Param			refresh_token	body		controllers.RefreshTokenController	true	"Refresh token"
//	@Success		200				{object}	server.Result[views.RefreshTokensResponse]
//	@Router			/users/{user_id}/tokens/refresh [post]
func (h *accessHandler) refreshTokens(c echo.Context) error {
	ctx := c.Request().Context()

	var req controllers.RefreshTokensRequest
	if err := c.Bind(&req); err != nil {
		err = fmt.Errorf("failed to decode request. Reason: %v", err)
		log.Println(err)

		return server.ReturnError(c, http.StatusInternalServerError, err)
	}

	accessToken, refreshToken, err := h.uc.RefreshTokensHandler.Execute(ctx, req.RefreshToken, c.RealIP())
	if err != nil {
		err := fmt.Errorf("failed to refresh tokens. Reason: %v", err)
		log.Println(err)

		return server.ReturnError(c, http.StatusInternalServerError, err)
	}

	return server.ReturnResult(c, views.NewRefreshTokensResponse(accessToken, refreshToken))
}

// getID godoc
//
//	@Summary		Get user id
//	@Description	Return user id by access and refresh tokens
//	@Tags			access
//	@Accept			json
//	@Produce		json
//	@Param			access_token	body		controllers.AccessTokenController	true	"Access token"
//	@Success		200				{object}	server.Result[views.GetIDResponse]
//	@Router			/users/id [post]
func (h *accessHandler) getID(c echo.Context) error {
	ctx := c.Request().Context()

	var req controllers.GetIDRequest
	if err := c.Bind(&req); err != nil {
		err = fmt.Errorf("failed to decode request. Reason: %v", err)
		log.Println(err)

		return server.ReturnError(c, http.StatusInternalServerError, err)
	}

	userID, err := h.uc.GetIDHandler.Execute(ctx, req.AccessToken)
	if err != nil {
		err := fmt.Errorf("failed to get id. Reason: %v", err)
		log.Println(err)

		return server.ReturnError(c, http.StatusInternalServerError, err)
	}

	return server.ReturnResult(c, views.NewGetIDResponse(userID))
}

// deauthoriseUser godoc
//
//	@Summary		Deauthorise user
//	@Description	Denies access for getting id and refresh tokens
//	@Tags			access
//	@Accept			json
//	@Produce		json
//	@Param			access_token	body		controllers.AccessTokenController	true	"Access token"
//	@Success		200				{object}	server.Result[views.DeauthoriseUserResponse]
//	@Router			/users/deauthorise [post]
func (h *accessHandler) deauthoriseUser(c echo.Context) error {
	ctx := c.Request().Context()

	var req controllers.DeauthoriseUserRequest
	if err := c.Bind(&req); err != nil {
		err = fmt.Errorf("failed to decode request. Reason: %v", err)
		log.Println(err)

		return server.ReturnError(c, http.StatusInternalServerError, err)
	}

	err := h.uc.DeauthoriseUserHandler.Execute(ctx, req.AccessToken)
	if err != nil {
		err := fmt.Errorf("failed to deauthorise user. Reason: %v", err)
		log.Println(err)

		return server.ReturnError(c, http.StatusInternalServerError, err)
	}

	return server.ReturnResult(c, views.NewDeauthoriseUserResponse())
}

func makeAccessRoutes(srv *server.Server, uc *usecases.UseCases) {
	sg := srv.BasePath()
	h := accessHandler{uc: uc}

	{
		sg := sg.Group("/users")

		sg.POST("", h.createUser)
		sg.POST("/id", h.getID)
		sg.POST("/deauthorise", h.deauthoriseUser)

		{
			sg = sg.Group("/:user_id/tokens")

			sg.POST("", h.getTokens)
			sg.POST("/refresh", h.refreshTokens)
		}

	}
}
