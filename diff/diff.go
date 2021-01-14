package diff

import (
	"fmt"
	"sort"
)

func PrintDiff(leftFileName, rightFileName string) {
	leftSourceType, err := detectSourceType(leftFileName)
	check(err)

	rightSourceType, err := detectSourceType(rightFileName)
	check(err)

	leftParse := parser(leftSourceType)
	left, err := leftParse(leftFileName)
	check(err)

	rightParse := parser(rightSourceType)
	right, err := rightParse(rightFileName)
	check(err)

	fmt.Println(leftSourceType, "vs", rightSourceType)

	findDiff("<", left, right)
	findDiff(">", right, left)
}

func findDiff(prefix string, left, right map[string][][]string) {
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
