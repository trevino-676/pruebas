package test

import (
	"testing"

	"github.com/trevino-676/pruebas/pkg"
)

func assertFizzBuzz(t *testing.T, got string, expected string) {
	if expected != got {
		t.Errorf("Error on test, got %s expected %s", got, expected)
	}
}

var fizzBuzzTest = []struct {
	name           string
	assertFunc     func(t *testing.T, got string, expected string)
	testCase       int
	expectedResult string
}{
	{name: "test_0", assertFunc: assertFizzBuzz, testCase: 0, expectedResult: "0"},
	{name: "test_anyNumber", assertFunc: assertFizzBuzz, testCase: 1, expectedResult: "1"},
	{name: "test_multiple_3", assertFunc: assertFizzBuzz, testCase: 3, expectedResult: "Fizz"},
	{name: "test_multiple_5", assertFunc: assertFizzBuzz, testCase: 5, expectedResult: "Buzz"},
	{name: "test_multiple_3_and_5", assertFunc: assertFizzBuzz, testCase: 15, expectedResult: "FizzBuzz"},
}

func TestFizzBuzz(t *testing.T) {
	for _, tt := range fizzBuzzTest{
		t.Run(tt.name, func(t *testing.T) {
			gotResult := pkg.FizzBuzz(tt.testCase)
			tt.assertFunc(t, gotResult, tt.expectedResult)
		})
	}
}
