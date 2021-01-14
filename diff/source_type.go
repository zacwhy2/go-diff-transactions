package diff

import (
	"bufio"
	"os"
	"strings"

	"github.com/zacwhy/go-diff-transactions/array"
)

func detectSourceType(fileName string) (string, error) {
	f, err := os.Open(fileName)

	if err != nil {
		return "", err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for i := 0; scanner.Scan(); i++ {
		text := scanner.Text()

		if i == 0 {
			if text == `"Transaction Date","Description","Amount"` {
				return "po", nil
			}

			headers := strings.Split(text, ",")
			dateIndex := array.IndexOf("d", headers)
			amountIndex := array.IndexOf("a", headers)
			descriptionIndex := array.IndexOf("c", headers)

			if dateIndex != -1 && amountIndex != -1 && descriptionIndex != -1 {
				return "local", nil
			}
		}

		if i == 4 && text == "Date,DESCRIPTION,Foreign Currency Amount,SGD Amount" {
			return "sc", nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", nil
}

func parser(sourceType string) func(string) (map[string][][]string, error) {
	switch sourceType {
	case "local":
		return ParseLocal
	case "po":
		return ParsePo
	case "sc":
		return ParseSc
	}
	return nil
}
