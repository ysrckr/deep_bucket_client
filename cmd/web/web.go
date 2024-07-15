package web

import (
	"embed"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	//go:embed dist
	dist embed.FS
)

func RegisterHandlers(e *echo.Echo) {

	if os.Getenv("ENV") == "development" {
		log.Println("Running in dev mode")
		setupDevProxy(e)
		return
	}

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:       "dist",
		Index:      "index.html",
		Browse:     false,
		HTML5:      true,
		Filesystem: http.FS(dist),
	}))

}

func setupDevProxy(e *echo.Echo) {
	url, err := url.Parse("http://localhost:5173")
	if err != nil {
		log.Fatal(err)
	}
	// Setep a proxy to the vite dev server on localhost:5173
	balancer := middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
		{
			URL: url,
		},
	})

	e.Use(middleware.ProxyWithConfig(middleware.ProxyConfig{
		Balancer: balancer,
		Skipper: func(c echo.Context) bool {
			// Skip the proxy if the prefix is /api
			return len(c.Path()) >= 4 && c.Path()[:4] == "/api"
		},
	}))
}
