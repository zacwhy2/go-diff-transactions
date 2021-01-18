package diff

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"github.com/zacwhy/go-diff-transactions/array"
)

type source interface {
	parse(fileName string) (recordGroups, error)
}

func detectSource(fileName string) (source, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for i := 0; scanner.Scan(); i++ {
		text := scanner.Text()

		if i == 0 {
			if text == `"Transaction Date","Description","Amount"` {
				return poSource{}, nil
			}

			headers := strings.Split(text, ",")
			dateIndex := array.IndexOf("d", headers)
			amountIndex := array.IndexOf("a", headers)
			descriptionIndex := array.IndexOf("c", headers)

			if dateIndex != -1 && amountIndex != -1 && descriptionIndex != -1 {
				return localSource{}, nil
			}
		}

		if i == 4 && text == "Date,DESCRIPTION,Foreign Currency Amount,SGD Amount" {
			return scSource{}, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return nil, errors.New("cannot detect source type for: " + fileName)
}
