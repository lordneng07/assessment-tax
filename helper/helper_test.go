//go:build unit
// +build unit

package helper

import "testing"

func TestHelper(t *testing.T) {
	t.Run("requird tag should be totalIncome", func(t *testing.T) {
		want := "totalIncome"
		got := CamelCase("TotalIncome")

		if want != got {
			t.Errorf("got = %s, want = %s", got, want)
		}
	})
}
