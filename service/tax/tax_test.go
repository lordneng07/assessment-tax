package service

import (
	"testing"
)

func TestTaxNet(t *testing.T) {

	cases := []struct {
		name      string
		income    float64
		wht       float64
		allowance float64
		tax       float64
	}{
		{
			name:      `not pay tax input income = 180000, wht = 0, allowance = 0, should be 0`,
			income:    180000.0,
			wht:       0.0,
			allowance: 0.0,
			tax:       0.0,
		},
		{
			name:      `not pay tax and got refund input income = 180000, wht = 9000, allowance = 0, should be 0`,
			income:    180000.0,
			wht:       9000.0,
			allowance: 0.0,
			tax:       0.0,
		},
		{
			name:      `not pay tax and got refund input income = 60000, wht = 3000, allowance = 0, should be 0`,
			income:    60000.0,
			wht:       3000.0,
			allowance: 0.0,
			tax:       0.0,
		},
		{
			name:      `tax rate 10% input income = 500000, wht = 0, allowance = 0, should be 29000`,
			income:    500000.0,
			wht:       0.0,
			allowance: 0.0,
			tax:       29000.0,
		},
		{
			name:      `tax rate 10% input income = 550000, wht = 0, allowance = 0, should be 34000`,
			income:    550000.0,
			wht:       0.0,
			allowance: 0.0,
			tax:       34000.0,
		},
		{
			name:      `rate 15% input income = 1000000, wht = 0, allowance = 0, should be 66000`,
			income:    1000000.0,
			wht:       0.0,
			allowance: 0.0,
			tax:       66000.0,
		},
		{
			name:      `rate 20% input income = 2000000, wht = 0 allowance = 0, should be 188000`,
			income:    2000000.0,
			wht:       0.0,
			allowance: 0.0,
			tax:       188000.0,
		},
		{
			name:      `rate 35% input income = 2100000, wht = 0, allowance = 0, should be 714000`,
			income:    2100000.0,
			wht:       0.0,
			allowance: 0.0,
			tax:       714000.0,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tn := NewTaxService(tt.income, tt.wht, tt.allowance).TaxNet()

			if tn != tt.tax {
				t.Errorf("got = %.2f, want = %.2f", tn, tt.tax)
			}
		})
	}
}

func TestRefund(t *testing.T) {
	cases := []struct {
		name      string
		income    float64
		wht       float64
		allowance float64
		refund    float64
	}{
		{
			name:      `not pay tax and got refund input income = 180000, wht = 9000, allowance = 0, should be 9000`,
			income:    180000.0,
			wht:       9000.0,
			allowance: 0.0,
			refund:    9000.0,
		},
		{
			name:      `not pay tax and got refund input income = 60000, wht = 3000, allowance = 0, should be 3000`,
			income:    60000.0,
			wht:       3000.0,
			allowance: 0.0,
			refund:    3000.0,
		},
		{
			name:      `not pay tax and got refund input income = 500000, wht = 39000, allowance = 0, should be 10000`,
			income:    500000.0,
			wht:       39000.0,
			allowance: 0.0,
			refund:    10000.0,
		},
		{
			name:      `not pay tax and got refund input income = 550000, wht = 54000, allowance = 0, should be 20000`,
			income:    550000.0,
			wht:       54000.0,
			allowance: 0.0,
			refund:    20000.0,
		},
		{
			name:      `not pay tax and got refund input income = 1000000, wht = 70000, allowance = 0, should be 4000`,
			income:    1000000.0,
			wht:       70000.0,
			allowance: 0.0,
			refund:    4000.0,
		},
		{
			name:      `not pay tax and got refund input income = 2000000, wht = 200000, allowance = 0, should be 12000`,
			income:    2000000.0,
			wht:       200000.0,
			allowance: 0.0,
			refund:    12000.0,
		},
		{
			name:      `not pay tax and got refund input income = 2100000, wht = 720000, allowance = 0, should be 6000`,
			income:    2100000.0,
			wht:       720000.0,
			allowance: 0.0,
			refund:    6000.0,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tn := NewTaxService(tt.income, tt.wht, tt.allowance).Refund()

			if tn != tt.refund {
				t.Errorf("got = %.2f, want = %.2f", tn, tt.refund)
			}
		})
	}
}
