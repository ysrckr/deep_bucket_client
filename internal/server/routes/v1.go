package routes

import "github.com/labstack/echo/v4"

func (r *Router) V1Routes() *echo.Group {

	v1 := r.ApiRoutes().Group("/v1")

	return v1
}
