package diff

import "testing"

func TestDiffLocalVsPo(t *testing.T) {
	localRecordGroups, err := localSource{}.parse("testdata/po/local.csv")

	if err != nil {
		t.Fatal(err)
	}

	poRecordGroups, err := poSource{}.parse("testdata/po/remote.csv")

	if err != nil {
		t.Fatal(err)
	}

	leftDiffKeys := findDiff(localRecordGroups, poRecordGroups)

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

	rightDiffKeys := findDiff(poRecordGroups, localRecordGroups)

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
