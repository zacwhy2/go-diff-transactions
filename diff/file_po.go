package diff

import (
	"encoding/csv"
	"io"
	"math/big"
	"os"
	"regexp"
	"time"
)

type poSource struct {
}

func (poSource poSource) String() string {
	return "po"
}

func (poSource poSource) parse(fileName string) (recordGroups, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	m := make(recordGroups)

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	re := regexp.MustCompile(`S\$(\d{1,}.\d{2})`)

	for i := 0; ; i++ {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		if i < 2 {
			continue
		}

		if record[0] == "Sub-Total" {
			break
		}

		time, err := time.Parse("02 Jan 2006", record[0])

		if err != nil {
			return nil, err
		}

		matches := re.FindStringSubmatch(record[2])

		dollars := new(big.Rat)
		dollars.SetString(matches[1])

		cents := dollars.Mul(dollars, big.NewRat(100, 1))

		key := time.Format("2006-01-02") + " " + cents.FloatString(0)

		if _, ok := m[key]; !ok {
			m[key] = [][]string{}
		}

		m[key] = append(m[key], record)
	}

	return m, nil
}
