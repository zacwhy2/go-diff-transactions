package diff

import (
	"errors"
	"fmt"
	"math/big"
	"sort"
	"strings"
)

type recordGroups map[string][][]string

// PrintDiff prints the difference between two sources
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
		s := strings.Split(key, " ")
		date := s[0]

		cents, ok := new(big.Rat).SetString(s[1])
		if !ok {
			check(errors.New("error with record amount"))
		}
		dollar := cents.Quo(cents, big.NewRat(100, 1))
		amount := dollar.FloatString(2)

		for i, row := range recordGroups[key] {
			fmt.Printf("%v %v $%v #%v %v\n", prefix, date, amount, i+1, row)
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
