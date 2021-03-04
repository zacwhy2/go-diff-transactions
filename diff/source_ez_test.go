package diff

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	want := recordGroups{
		"2020-01-01 100": [][]string{
			[]string{"01/01/2020 19:10:05", "-1.00", "BOTH"},
		},
		"2020-01-02 100": [][]string{
			[]string{"02/01/2020 19:46:32", "-1.00", "EZ ONLY"},
		},
	}

	got, err := ezSource{}.parse("testdata/ez.csv")

	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("parse() mismatch (-want +got):\n%s", diff)
	}
}
