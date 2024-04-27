package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

type (
	TaxHandler interface {
		Calculation(c echo.Context) error
	}

	taxHendler struct {
	}

	TaxRequest struct {
		Income     float64         `json:"totalIncome"`
		Wht        float64         `json:"wht"`
		Allowances []AllowanceType `json:"allowances"`
	}

	TaxResponse struct {
		Tax float64 `json:"tax"`
	}

	AllowanceType struct {
		AllowanceType string  `json:"allowanceType"`
		Amount        float64 `json:"amount"`
	}
)

func (h taxHendler) Calculation(c echo.Context) (err error) {
	req := TaxRequest{}

	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	return c.JSON(http.StatusOK, req)
}
