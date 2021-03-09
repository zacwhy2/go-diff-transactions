package xls

import (
	"testing"

	"github.com/extrame/xls"
	"github.com/google/go-cmp/cmp"
)

func TestReadRows(t *testing.T) {
	want := [][]string{
		[]string{"a1", "b1", ""},
		[]string{"a2", "b2", ""},
	}

	workBook, err := xls.Open("testdata/sample.xls", "utf-8")

	if err != nil {
		t.Fatal(err)
	}

	sheet := workBook.GetSheet(0)

	if sheet == nil {
		t.Fatal("no worksheet")
	}

	got := ReadRows(sheet)

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("ReadRows(sheet) mismatch (-want +got):\n%s", diff)
	}
}
