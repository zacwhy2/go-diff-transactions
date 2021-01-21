package diff

import "testing"

func TestFindDiff(t *testing.T) {
	a := make(recordGroups)
	b := make(recordGroups)

	var diffKeys []string

	diffKeys = findDiff(a, b)

	{
		want := 0
		got := len(diffKeys)
		if got != want {
			t.Errorf("len(diffKeys) = %v; want %v", got, want)
		}
	}

	a["key1"] = [][]string{}

	diffKeys = findDiff(a, b)

	{
		want := 1
		got := len(diffKeys)
		if got != want {
			t.Errorf("len(diffKeys) = %d; want %d", got, want)
		}
	}

	{
		want := "key1"
		got := diffKeys[0]
		if got != want {
			t.Errorf("diffKeys[0] = %q; want %q", got, want)
		}
	}
}
