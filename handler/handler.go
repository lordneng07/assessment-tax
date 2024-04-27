package handler

import (
	"net/http"

	"github.com/labstack/echo"
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

	g.POST("tax/calculations", h.TaxHandler.Calculation)
}

func Echo() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Go Bootcamp!")
	})

	return e
}
