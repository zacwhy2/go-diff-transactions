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

	fmt.Println(leftSource)
	leftDiffKeys := findDiff(left, right)
	printRecords("<", leftDiffKeys, left)

	fmt.Println(rightSource)
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

	return keys
}

func printRecords(prefix string, keys []string, recordGroups recordGroups) {
	sort.Strings(keys)
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
