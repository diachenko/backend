package compute

import (
	"math"
	"strconv"
	"testing"
)

var exps = map[string]float64{
	"1+1":         2,
	"1+2^3^2":     513,
	"2^(3+4)":     128,
	"2^(3/(1+2))": 2,
	"2(1+3)":      8,
	"2^2(1+3)":    16,
	"-2":          -2,
	"1+(-1)^2":    2,
	"3-5":         -2,
	"-3*4":        -12,
	"3*-4":        -12,
	"3*(3-(5+6)^12)*23^3-5^23": -126476703133661843,
	"2^3^2":                    512,
	"-3^2":                     -9,
	"2(1+1)4":                  16,
}

const DELTA = 0.000001

func TestEvaluate(t *testing.T) {
	for expression, expected := range exps {
		res, err := Evaluate(expression)
		if err != nil {
			t.Error(err)
		} else if math.Abs(res-expected) > DELTA {
			message := expression + " failed: actual value " +
				strconv.FormatFloat(res, 'G', -1, 64) +
				" differs from expected value " +
				strconv.FormatFloat(expected, 'G', -1, 64)
			t.Error(message)
		}
	}
}

func BenchmarkEvaluate(b *testing.B) {
	tests := []string{
		"1+2^3^2",
		"2^(3+4)",
		"2^(3/(1+2))",
		"2^2(1+3)",
		"1+(-1)^2",
		"3*(3-(5+6)^12)*23^3-5^23",
		"2^3^2",
	}
	for i := 0; i < b.N; i++ {
		Evaluate(tests[i%len(tests)])
	}
}
