package handler

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
	service "github.com/lordneng07/assessment-tax/service/tax"
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
		Tax    float64 `json:"tax"`
		Refund float64 `json:"taxRefund"`
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

	isError, verr := req.validate()

	if isError {
		return c.JSON(http.StatusBadRequest, verr)
	}

	return c.JSON(http.StatusOK, NewTaxResponse(req))
}

func NewTaxResponse(tr TaxRequest) TaxResponse {

	return TaxResponse{
		Tax:    service.NewTaxService(tr.Income, tr.Wht, tr.Allowance()).TaxNet(),
		Refund: service.NewTaxService(tr.Income, tr.Wht, tr.Allowance()).Refund(),
	}
}

func (tr TaxRequest) validate() (bool, Err) {

	for _, al := range tr.Allowances {
		if !strings.Contains(al.AllowanceType, "donation") && !strings.Contains(al.AllowanceType, "k-receipt") {
			return true, Err{
				Message: "allowanceType is not correct",
			}
		}
	}

	return false, Err{
		Message: "Success",
	}
}

func (tr TaxRequest) Allowance() float64 {
	alw := 0.0
	for _, al := range tr.Allowances {

		alw += al.Amount
	}

	return alw
}
