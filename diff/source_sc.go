package diff

import (
	"encoding/csv"
	"io"
	"math/big"
	"os"
	"regexp"
	"time"
)

type scSource struct {
}

func (scSource scSource) String() string {
	return "sc"
}

func (scSource scSource) parse(fileName string) (recordGroups, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	m := make(recordGroups)

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	reader.TrimLeadingSpace = true

	re := regexp.MustCompile(`SGD (\d{1,}.\d{2}) (CR|DR)`)

	for i := 0; ; i++ {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		if i < 4 {
			continue
		}

		// TODO
		if record[0] == "Current Balance" {
			break
		}

		if record[0] == "" {
			continue // skip blank lines
		}

		time, err := time.Parse("02/01/2006", record[0])

		if err != nil {
			return nil, err
		}

		matches := re.FindStringSubmatch(record[3])

		if matches[2] != "DR" {
			continue // skip non-DR lines
		}

		amount := new(big.Rat)
		amount.SetString(matches[1])

		cents := amount.Mul(amount, big.NewRat(100, 1))

		key := time.Format("2006-01-02") + " " + cents.FloatString(0)

		if _, ok := m[key]; !ok {
			m[key] = [][]string{}
		}

		m[key] = append(m[key], record)
	}

	return m, nil
}
