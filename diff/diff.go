package diff

import (
	"fmt"
	"sort"
)

type recordGroups map[string][][]string

func PrintDiff(leftFileName, rightFileName string) {
	leftSource, err := detectSource(leftFileName)
	check(err)

	rightSource, err := detectSource(rightFileName)
	check(err)

	left, err := leftSource.parse(leftFileName)
	check(err)

	right, err := rightSource.parse(rightFileName)
	check(err)

	fmt.Println(leftSource, "<>", rightSource)

	findDiff("<", left, right)
	findDiff(">", right, left)
}

func findDiff(prefix string, left, right recordGroups) {
	var keys []string
	for k := range left {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		leftRows := left[key]
		rightRows, ok := right[key]
		if !ok || len(leftRows) != len(rightRows) {
			for i, row := range leftRows {
				fmt.Println(prefix, key, i, row)
			}
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
