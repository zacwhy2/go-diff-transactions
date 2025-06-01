package diff

import (
	"encoding/csv"
	"io"
	"math/big"
	"os"
	"regexp"
	"time"
)

type hsSource struct {
}

func (hsSource hsSource) String() string {
	return "hs"
}

func (hsSource hsSource) parse(fileName string) (recordGroups, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	m := make(recordGroups)

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	reDate := regexp.MustCompile(`([0-9]{2}/[0-9]{2}/[0-9]{4}) •••• •••• •••• [0-9]{4} -[0-9]{1,}\.[0-9]{2} SGD$`)
	reAmount := regexp.MustCompile(`^-([0-9]{1,}\.[0-9]{2})`)

	for i := 0; ; i++ {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		// Date
		matchesDate := reDate.FindStringSubmatch(record[1])
		if len(matchesDate) == 0 {
			continue
		}
		time, err := time.Parse("02/01/2006", matchesDate[1])
		if err != nil {
			return nil, err
		}

		// Amount
		matches := reAmount.FindStringSubmatch(record[2])
		if len(matches) == 0 {
			continue
		}
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
