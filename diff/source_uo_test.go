package diff

import (
	"testing"
)

func TestParseArray(t *testing.T) {
	cells := [][]string{
		[]string{"Lorem Ipsum"},
		[]string{},
		[]string{"Account Statement Details"},
		[]string{},
		[]string{"Account Number:", "1234567890123456", "SGD"},
		[]string{"Account Type:", ""},
		[]string{"Statement Date:", "01 Feb 2020"},
		[]string{"Statement Balance:", "2.00", "SGD"},
		[]string{},
		[]string{
			"Transaction Date",
			"Posting Date",
			"Description",
			"Foreign Currency Type",
			"Transaction Amount(Foreign)",
			"Local Currency Type",
			"Transaction Amount(Local)",
		},
		[]string{},
		[]string{"01 Jan 2020", "01 Jan 2020", "both", "", "", "SGD", "1.00"},
		[]string{"02 Jan 2020", "02 Jan 2020", "uo only", "", "", "SGD", "1.00"},
	}

	want := map[string][][]string{
		"2020-01-01 100": [][]string{
			[]string{"01 Jan 2020", "1.00", "both"},
		},
		"2020-01-02 100": [][]string{
			[]string{"02 Jan 2020", "1.00", "uo only"},
		},
	}

	got, err := parseArray(cells)

	if err != nil {
		t.Errorf("err = %v; want %v", err, nil)
	}

	if !groupsEqual(got, want) {
		t.Errorf("parseArray(cells) = %v; want %v", got, want)
	}
}

func groupsEqual(left, right map[string][][]string) bool {
	if len(left) != len(right) {
		return false
	}

	for leftKey, leftValue := range left {
		rightValue, ok := right[leftKey]

		if !ok {
			return false
		}

		if len(leftValue) != len(rightValue) {
			return false
		}

		for i := 0; i < len(leftValue); i++ {
			for j := 0; j < len(leftValue[i]); j++ {
				if leftValue[i][j] != rightValue[i][j] {
					return false
				}
			}
		}
	}

	return true
}
