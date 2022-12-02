package calculator_test

import (
	"fmt"
	"github.com/ReyAdrian520/PivotSchool/calculator"
	"testing"
)

func TestCalculator(t *testing.T) {
	testCases := []struct {
		a        int
		b        int
		operator string
		want     int
	}{
		{1000, 400, "+", 1400},
		{444, 444, "+", 888},
		{65, 3, "+", 68},
		{498, 455, "-", -43},
		{12, 6, "-", -6},
		{6, 12, "-", 6},
		{100, 100, "*", 10000},
		{7987, 0, "*", 0},
		{102, 3, "*", 306},
		{22, 11, "/", 2},
		{144, 12, "/", 12},
		{6754, 0, "/", 0},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d%s%d", tc.a, tc.operator, tc.b), func(t *testing.T) {
			switch tc.operator {
			case "+":
				if got := calculator.Add(tc.a, tc.b); got != tc.want {
					t.Errorf("got: %d - want: %d", got, tc.want)
				}
			case "-":
				if got := calculator.Sub(tc.a, tc.b); got != tc.want {
					t.Errorf("got: %d - want: %d", got, tc.want)
				}
			case "*":
				if got := calculator.Muti(tc.a, tc.b); got != tc.want {
					t.Errorf("got: %d - want: %d", got, tc.want)
				}
			case "/":
				if got, err := calculator.Div(tc.a, tc.b); err != nil {
					if tc.b != 0 {
						t.Errorf("%d is not 0 but returned error", tc.b)
					}
				} else if got != tc.want {
					t.Errorf("got: %d - want: %d", got, tc.want)
				}
			default:
				t.Errorf("invalid operator: %s", tc.operator)
			}
		})
	}
}
