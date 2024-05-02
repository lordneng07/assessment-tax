package handler

import (
	"net/http"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/labstack/echo"
	"github.com/lordneng07/assessment-tax/helper"
	"github.com/lordneng07/assessment-tax/model"
	"github.com/lordneng07/assessment-tax/service"
)

type (
	TaxHandler interface {
		Calculate(c echo.Context) error
	}

	taxHendler struct {
		service.TaxLevelService
	}

	TaxRequest struct {
		TotalIncome *float64        `json:"totalIncome" validate:"required,gt=0"`
		Wht         *float64        `json:"wht" validate:"required,number,gte=0" errormgs:"Invalid wht is required"`
		Allowances  []AllowanceType `json:"allowances" validate:"required"`
	}

	AllowanceType struct {
		AllowanceType string  `json:"allowanceType"`
		Amount        float64 `json:"amount"`
	}

	Donation struct {
		Calculator
	}
)

var uni *ut.UniversalTranslator

func (h taxHendler) Calculation(c echo.Context) (err error) {
	req := TaxRequest{}

	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	en := en.New()
	uni = ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	validate := validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)
	errValidate := validate.Struct(req)

	if errValidate != nil {
		errs := translateError(errValidate, trans)
		return c.JSON(http.StatusBadRequest, errs)
	}

	isError, er := req.validate()

	if isError {
		return c.JSON(http.StatusBadRequest, er)
	}

	return c.JSON(http.StatusOK, h.NewTax(*req.TotalIncome, *req.Wht, req.Allowance()))
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
	return tr.donation()
}

func (tr TaxRequest) donation() float64 {
	donate := 0.0
	for _, al := range tr.Allowances {

		if strings.Contains(al.AllowanceType, "donation") {
			donate += al.Amount
		}
	}

	if donate > 100000 {
		return 100000
	}

	return donate
}

func translateError(err error, trans ut.Translator) (errs []model.ApiErrorResponse) {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	out := make([]model.ApiErrorResponse, len(validatorErrs))
	for i, e := range validatorErrs {
		translatedErr := e.Translate(trans)
		translatedErr = strings.ReplaceAll(translatedErr, e.Field(), helper.CamelCase(e.Field()))
		out[i] = model.ApiErrorResponse{Field: helper.CamelCase(e.Field()), Msg: translatedErr}
	}
	return out
}
