package diff

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/zacwhy/go-diff-transactions/array"
)

type localSource struct {
}

func (localSource localSource) String() string {
	return "local"
}

func (localSource localSource) parse(fileName string) (recordGroups, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	m := make(recordGroups)
	reader := csv.NewReader(file)

	var dateIndex, amountIndex int

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

		if _, ok := m[key]; !ok {
			m[key] = [][]string{}
		}

		m[key] = append(m[key], record)
	}

	return m, nil
}
