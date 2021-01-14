package diff

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/zacwhy/go-diff-transactions/array"
)

func ParseLocal(fileName string) (map[string][][]string, error) {
	f, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	m := make(map[string][][]string)

	reader := csv.NewReader(f)

	var dateIndex int
	var amountIndex int

	for i := 0; ; i++ {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		if i == 0 {
			dateIndex = array.IndexOf("d", record)
			amountIndex = array.IndexOf("a", record)
			continue
		}

		date := record[dateIndex]
		amount := record[amountIndex]
		key := date + " " + amount

		_, ok := m[key]

		if !ok {
			m[key] = [][]string{}
		}

		m[key] = append(m[key], record)
	}

	return m, nil
}
