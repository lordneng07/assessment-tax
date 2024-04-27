package service

type (
	TaxService interface {
		Calculate()
	}

	taxService struct {
		income    float64
		wht       float64
		allowance float64
	}
)

func NewTaxService(income float64, wht float64, allowance float64) taxService {
	return taxService{
		income:    income,
		wht:       wht,
		allowance: allowance,
	}
}

func (t taxService) Calculate() float64 {

	if t.netIncome() >= 150001 && t.netIncome() <= 500000 {
		return t.level(150000.0, 0.10) - t.wht
	}

	if t.netIncome() >= 500001 && t.netIncome() <= 1000000 {
		return t.level(500000.0, 0.15) - t.wht
	}

	if t.netIncome() >= 1000001 && t.netIncome() <= 2000000 {
		return t.level(1000000.0, 0.20) - t.wht
	}

	if t.netIncome() > 2000000 {
		return t.level(0, 0.35) - t.wht
	}

	return 0 - t.wht
}

func (t taxService) netIncome() float64 {
	return t.income - 60000 - t.allowance
}

func (t taxService) level(mpr float64, r float64) float64 {
	return (t.netIncome() - mpr) * r
}
