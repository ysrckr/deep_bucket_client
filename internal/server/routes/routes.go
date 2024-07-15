package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ysrckr/deep_bucket_client/cmd/web"
)

type Router struct {
	router *echo.Echo
}

func NewRouter() *Router {
	return &Router{
		router: echo.New(),
	}
}

func (r *Router) RegisterRoutes() http.Handler {

	r.router.Pre(middleware.RemoveTrailingSlash())
	r.router.Use(middleware.Logger())
	r.router.Use(middleware.Recover())

	// Setup the frontend handlers to service vite static assets

	web.RegisterHandlers(r.router)

	r.ApiRoutes()
	r.V1Routes()

	return r.router
}
