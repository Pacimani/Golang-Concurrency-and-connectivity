package sqrt

import (
	"fmt"
	"testing"
)

func almostEqual(a, b float64) bool {
	return Abs(a-b) < 0.001
}

func TestSimple(t *testing.T) {

	val, err := Sqrt(2)

	if err != nil {
		t.Errorf("got error %v", err)
	}
	if !almostEqual(val, 1.414214) {
		t.Errorf("Got %v, want %v", val, 1.414213)
	}
}

type testCase struct {
	x    float64
	want float64
}

func TestMany(t *testing.T) {

	tests := []testCase{
		{0.0, 0.0},
		{1.0, 1.0},
		{2.0, 1.414213},
		{3.0, 1.732051},
		{4.0, 2.0},
		{5.0, 2.236068},
		{6.0, 2.449489},
		{7.0, 2.645751},
	}

	for _, tc := range tests {

		t.Run(fmt.Sprintf("Sqrt(%v)", tc.x), func(t *testing.T) {

			val, err := Sqrt(tc.x)
			if err != nil {
				t.Errorf("got error %v", err)
			}
			if !almostEqual(val, tc.want) {
				t.Errorf("Got %v, want %v", val, tc.want)
			}
		})

	}
}
