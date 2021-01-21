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

	leftDiffKeys := findDiff(left, right)
	printRecords("<", leftDiffKeys, left)

	rightDiffKeys := findDiff(right, left)
	printRecords(">", rightDiffKeys, right)
}

func findDiff(left, right recordGroups) []string {
	var keys []string

	for key, leftRows := range left {
		rightRows, ok := right[key]
		if !ok || len(leftRows) != len(rightRows) {
			keys = append(keys, key)
		}
	}

	sort.Strings(keys)
	return keys
}

func printRecords(prefix string, keys []string, recordGroups recordGroups) {
	for _, key := range keys {
		for i, row := range recordGroups[key] {
			fmt.Printf("%v [%v][%v] %v\n", prefix, key, i, row)
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
