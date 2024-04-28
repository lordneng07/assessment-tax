package handler

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Handler struct {
	TaxHandler
}

type Err struct {
	Message string `json:"message"`
}

func New() *Handler {
	return &Handler{
		TaxHandler: &taxHendler{},
	}
}

func SetApi(e *echo.Echo, h *Handler) {
	g := e.Group("/api/")
	g.Use(middleware.BasicAuth(authen))

	g.POST("tax/calculations", h.TaxHandler.Calculation)
}

func Echo() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Go Bootcamp!")
	})

	return e
}

func authen(username, password string, c echo.Context) (bool, error) {
	if username == os.Getenv("ADMIN_USERNAME") && password == os.Getenv("ADMIN_PASSWORD") {
		return true, nil
	}
	return false, nil
}
