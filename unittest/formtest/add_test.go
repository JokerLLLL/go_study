package formtest

import (
	"math"
	"testing"
)

func add(a,b int) int  {
	return a + b
}

func TestAdd(t *testing.T)  {
	tests := []struct{a,b,c int}{
		{1,2,3},
		{4,7,11},
		{-2, 0, -2},
		{0, 0, 0},
		{math.MaxInt64, 1, math.MinInt64},
	}
	for _, test := range tests {
		if actual := add(test.a, test.b); actual != test.c {
			t.Errorf("add a:%d and b:%d not eq c:%d actual:%d", test.a, test.b, test.c, actual)
		}
	}
}
