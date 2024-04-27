package service

import (
	"testing"
)

func TestCalculate(t *testing.T) {

	cases := []struct {
		name      string
		income    float64
		wht       float64
		allowance float64
		rate      float64
		tax       float64
	}{
		{
			name:      `not pay tax input income = 180000 wht = 0 allowance = 0 should be 0`,
			income:    180000.0,
			wht:       0.0,
			allowance: 0.0,
			rate:      0.0,
			tax:       0.0,
		},
		{
			name:      `not pay tax and got refund input income = 180000 wht = 9000 allowance = 0 should be 0`,
			income:    180000.0,
			wht:       9000.0,
			allowance: 0.0,
			rate:      0.0,
			tax:       -9000.0,
		},
		{
			name:      `not pay tax and got refund input income = 60000 wht = 3000 allowance = 0 should be 0`,
			income:    60000.0,
			wht:       3000.0,
			allowance: 0.0,
			rate:      0.0,
			tax:       -3000.0,
		},
		{
			name:      `tax rate 10% input income = 500000 wht = 0 allowance = 0 should be 29000`,
			income:    500000.0,
			wht:       0.0,
			allowance: 0.0,
			rate:      0.10,
			tax:       29000.0,
		},
		{
			name:      `tax rate 10% input income = 550000 wht = 0 allowance = 0 should be 34000`,
			income:    550000.0,
			wht:       0.0,
			allowance: 0.0,
			rate:      0.10,
			tax:       34000.0,
		},
		{
			name:      `rate 15% input income = 1000000 wht = 0 allowance = 0 should be 66000`,
			income:    1000000.0,
			wht:       0.0,
			allowance: 0.0,
			rate:      0.10,
			tax:       66000.0,
		},
		{
			name:      `rate 20% input income = 2000000 wht = 0 allowance = 0 should be 188000`,
			income:    2000000.0,
			wht:       0.0,
			allowance: 0.0,
			rate:      0.10,
			tax:       188000.0,
		},
		{
			name:      `rate 35% input income = 2100000 wht = 0 allowance = 0 should be 714000`,
			income:    2100000.0,
			wht:       0.0,
			allowance: 0.0,
			rate:      0.10,
			tax:       714000.0,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tn := NewTaxService(tt.income, tt.wht, tt.allowance).Calculate()

			if tn != tt.tax {
				t.Errorf("got = %.2f, want = %.2f", tn, tt.tax)
			}
		})
	}
}
