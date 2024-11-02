package sqrt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func almostEqual(a, b float64) bool {
	return Abs(a-b) < 0.001
}

func TestSimple(t *testing.T) {

	val, err := Sqrt(2)

	require.NoError(t, err)
	require.InDelta(t, 1.414214, val, 0.001)
	//require.Equal(t, 1.414214, val)
}
