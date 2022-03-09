package test

import (
	"testing"
	
	"github.com/trevino-676/pruebas/pkg"
)

type unitTest struct {
	name string
	assertFunc func(t *testing.T, got bool, want bool)
	expression string
	expectedResult bool
}

func assertBalanceExpresion(t *testing.T, got bool, want bool) {
	if got != want {
		t.Errorf("Error on test, expecte: %v want %v", want, got)
	}
}

var balanceExpresionTest = []unitTest{
	{
		name: "CorrectBalanceExpresion",
		assertFunc: assertBalanceExpresion,
		expression: "{[a * b (c + d)] - 5}",
		expectedResult: true,
	},
	{
		name: "IncorrectBalanceExpresion",
		assertFunc: assertBalanceExpresion,
		expression: "{ a * (c + d)] - 5}",
		expectedResult: false,
	},
}


func TestBalanceExpresion(t *testing.T){
	for _, tt := range balanceExpresionTest {
		t.Run(tt.name, func(t *testing.T) {
			result := pkg.BalanceExpresion(tt.expression)
			tt.assertFunc(t, result, tt.expectedResult)
		})
	}
}