package main

import (
	"fmt"
	"os"

	"github.com/zacwhy/go-diff-transactions/diff"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: " + os.Args[0] + " FILE1 FILE2")
		os.Exit(1)
	}

	leftFileName := os.Args[1]
	rightFileName := os.Args[2]

	diff.PrintDiff(leftFileName, rightFileName)
}
