package service

import "github.com/lordneng07/assessment-tax/model"

const PersonalDeduction = 60000

type Pay interface {
	TotalTaxLevel(income float64, wht float64) float64
}

type TaxLevelService interface {
	NewTax(totoalIncome, wht, allowance float64) model.Tax
}

type taxLevelService struct {
}

type payer struct {
	income    float64
	wht       float64
	allowance float64
}

type noPay struct {
}

type payTenPercent struct {
}

type payFifteenPercent struct {
}

type payTwentyPercent struct {
}

type payThirtyFivePercent struct {
}

func CreateNoPay() noPay {
	return noPay{}
}

func (np noPay) TotalTaxLevel(i, wht float64) float64 {
	if i <= 150000 {
		return 0
	}
	return 0
}

func CreatePayTenPercent() payTenPercent {
	return payTenPercent{}
}

func (pt payTenPercent) TotalTaxLevel(i float64, wht float64) float64 {
	if i >= 150001 && i <= 500000 {
		return Level(i, 150000.0, 0.10)
	}
	return 0
}

func CreatePayFifteenPercent() payFifteenPercent {
	return payFifteenPercent{}
}

func (pf payFifteenPercent) TotalTaxLevel(i float64, wht float64) float64 {
	if i >= 500001 && i <= 1000000 {
		return Level(i, 500000.0, 0.15)
	}
	return 0
}

func CreatePayTwentyPercent() payTwentyPercent {
	return payTwentyPercent{}
}

func (ptw payTwentyPercent) TotalTaxLevel(i float64, wht float64) float64 {
	if i >= 1000001 && i <= 2000000 {
		return Level(i, 1000000.0, 0.20)
	}

	return 0
}

func CreataPayThirtyFivePercent() payThirtyFivePercent {
	return payThirtyFivePercent{}
}

func (ptf payThirtyFivePercent) TotalTaxLevel(i float64, wht float64) float64 {
	if i >= 2000000 {
		return Level(i, 0, 0.35)
	}
	return 0
}

func (tls taxLevelService) NewTax(income, wht, allowance float64) model.Tax {
	payer := payer{
		income:    income,
		wht:       wht,
		allowance: allowance,
	}
	tl := []model.TaxLevel{}

	tl = append(tl, NewTaxLevel("0-150,000", payer.NetIncome(), payer.wht, CreateNoPay()))
	tl = append(tl, NewTaxLevel("150,001 - 500,000", payer.NetIncome(), payer.wht, CreatePayTenPercent()))
	tl = append(tl, NewTaxLevel("500,001 - 1,000,000", payer.NetIncome(), payer.wht, CreatePayFifteenPercent()))
	tl = append(tl, NewTaxLevel("1,000,001 - 2,000,000", payer.NetIncome(), payer.wht, CreatePayTwentyPercent()))
	tl = append(tl, NewTaxLevel("2,000,001 ขึ้นไป", payer.NetIncome(), payer.wht, CreataPayThirtyFivePercent()))

	taxLevel := getTaxLevel(tl)

	if taxLevel != nil {
		return model.Tax{
			Tax:      TotalTax(taxLevel.Tax, payer.wht),
			Refund:   TotalRefund(taxLevel.Tax, payer.wht),
			TaxLevel: tl,
		}
	}

	return model.Tax{
		Tax:      0,
		Refund:   TotalRefund(0, payer.wht),
		TaxLevel: tl,
	}
}

func NewTaxLevel(level string, i float64, wht float64, pay Pay) model.TaxLevel {
	return model.TaxLevel{
		Level: level,
		Tax:   pay.TotalTaxLevel(i, wht),
	}
}

func getTaxLevel(tl []model.TaxLevel) *model.TaxLevel {
	for _, nt := range tl {
		if nt.Tax > 0 {
			return &nt
		}
	}
	return nil
}

func (p *payer) NetIncome() float64 {
	return p.income - PersonalDeduction - p.allowance
}

func Level(i, mpr, r float64) float64 {
	return (i - mpr) * r
}

func TotalRefund(totalTax, wht float64) float64 {
	if wht > totalTax {
		return wht - totalTax
	}
	return 0
}

func TotalTax(nt, wht float64) float64 {
	if nt > wht {
		return nt - wht
	}

	return 0
}
