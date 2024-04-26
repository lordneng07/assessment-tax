package tax

func Calculate(totalIncome float64, wht float64, allowance float64) float64 {
	return calculateTax(calculateIncome(totalIncome, allowance), wht)
}

func calculateIncome(t float64, a float64) float64 {
	return t - 60000 - a
}

func calculateTax(i float64, wht float64) float64 {

	if i >= 150001 && i <= 500000 {
		return calculateTaxLevel(i, 150000.0, 0.10) - wht
	}

	if i >= 500001 && i <= 1000000 {
		return calculateTaxLevel(i, 500000.0, 0.15) - wht
	}

	if i >= 1000001 && i <= 2000000 {
		return calculateTaxLevel(i, 1000000.0, 0.20) - wht
	}

	if i > 2000000 {
		return calculateTaxLevel(i, 0, 0.35) - wht
	}

	return 0 - wht
}

func calculateTaxLevel(incomeNet float64, maxPreRate float64, ratio float64) float64 {
	return (incomeNet - maxPreRate) * ratio
}
