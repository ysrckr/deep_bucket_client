package routes

import "github.com/labstack/echo/v4"

func (r *Router) ApiRoutes() *echo.Group {

	api := r.router.Group("/api")

	return api
}
