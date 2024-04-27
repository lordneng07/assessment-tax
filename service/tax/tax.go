package service

const PersonalDeduction = 60000

type (
	Calculator interface {
		Calculate()
	}

	TaxService interface {
		Calculator
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

	if t.netIncome() <= 150000 {
		return 0
	}

	if t.netIncome() >= 150001 && t.netIncome() <= 500000 {
		return t.level(150000.0, 0.10)
	}

	if t.netIncome() >= 500001 && t.netIncome() <= 1000000 {
		return t.level(500000.0, 0.15)
	}

	if t.netIncome() >= 1000001 && t.netIncome() <= 2000000 {
		return t.level(1000000.0, 0.20)
	}

	if t.netIncome() > 2000000 {
		return t.level(0, 0.35)
	}

	return 0
}

func (t taxService) netIncome() float64 {
	return t.income - PersonalDeduction - t.allowance
}

func (t taxService) level(mpr float64, r float64) float64 {
	return ((t.netIncome() - mpr) * r)
}

func (t taxService) TaxNet() float64 {
	if t.Calculate() > t.wht {
		return t.Calculate() - t.wht
	}

	return 0
}

func (t taxService) Refund() float64 {

	if t.wht > t.Calculate() {
		return t.wht - t.Calculate()
	}

	return 0
}
