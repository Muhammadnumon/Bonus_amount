package main

import "testing"

func Test_bonusAmount(t *testing.T) {
	tests := []struct {
		name       string
		bonusSales [] int
		want       int
	}{
		{name: "Bonus", bonusSales: []int{9_000, 12_000, 16_000, 7_000}, want: 400},
		{name: "No Bonus", bonusSales: [] int{10_000}, want: 0},
	}
	for _, test := range tests {
		got := bonusAmount(test.bonusSales)
		if got != test.want {
			t.Errorf("bonusAmount() = %v, want %v", got, test.want)
		}

	}
}
