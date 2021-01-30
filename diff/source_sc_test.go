package diff

import "testing"

func TestDiffScVsLocal(t *testing.T) {
	scRecordGroups, err := scSource{}.parse("testdata/sc.csv")

	if err != nil {
		t.Fatal(err)
	}

	localRecordGroups, err := localSource{}.parse("testdata/local.csv")

	if err != nil {
		t.Fatal(err)
	}

	leftDiffKeys := findDiff(scRecordGroups, localRecordGroups)

	{
		want := 1
		got := len(leftDiffKeys)
		if got != want {
			t.Fatalf("len(leftDiffKeys) = %v; want %v", got, want)
		}
	}

	{
		want := "2020-01-02 100"
		got := leftDiffKeys[0]
		if got != want {
			t.Fatalf("leftDiffKeys[0] = %v; want %v", got, want)
		}
	}

	rightDiffKeys := findDiff(localRecordGroups, scRecordGroups)

	{
		want := 1
		got := len(rightDiffKeys)
		if got != want {
			t.Fatalf("len(rightDiffKeys) = %v; want %v", got, want)
		}
	}

	{
		want := "2020-01-03 100"
		got := rightDiffKeys[0]
		if got != want {
			t.Fatalf("rightDiffKeys[0] = %v; want %v", got, want)
		}
	}
}
