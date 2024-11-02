package sqrt

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

type testCase struct {
	val      float64
	expected float64
}

func loadSqrtCases(t *testing.T) []testCase {

	file, err := os.Open("stqt_case.csv")
	require.NoError(t, err)
	defer file.Close()

	var cases []testCase
	rdr := csv.NewReader(file)

	for {
		record, err := rdr.Read()
		if err == io.EOF {
			break
		}
		require.NoError(t, err)
		val, err := strconv.ParseFloat(record[0], 64)
		require.NoError(t, err)
		expected, err := strconv.ParseFloat(record[1], 64)
		require.NoError(t, err)
		cases = append(cases, testCase{val, expected})
	}
	return cases
}
func TestMany(t *testing.T) {

	cases := loadSqrtCases(t)
	for _, tc := range cases {
		t.Run(fmt.Sprintf("Sqrt(%v)", tc.val), func(t *testing.T) {
			out, err := Sqrt(tc.val)
			require.NoError(t, err)
			require.InDelta(t, tc.expected, out, 0.001)
		})
	}

}
