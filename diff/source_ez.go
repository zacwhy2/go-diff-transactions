package diff

import (
	"encoding/csv"
	"io"
	"math/big"
	"os"
	"time"
)

type ezSource struct {
}

func (s ezSource) String() string {
	return "ez"
}

func (s ezSource) parse(fileName string) (recordGroups, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	groups := make(recordGroups)
	reader := csv.NewReader(file)

	for i := 0; ; i++ {
		cells, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		if i < 1 {
			// headers
			continue
		}

		time, err := time.Parse("02/01/2006 15:04:05", cells[0])

		if err != nil {
			return nil, err
		}

		dollars := new(big.Rat)
		dollars.SetString(cells[1])
		cents := dollars.Mul(dollars, big.NewRat(-100, 1))

		key := time.Format("2006-01-02") + " " + cents.FloatString(0)

		if _, ok := groups[key]; !ok {
			groups[key] = [][]string{}
		}

		groups[key] = append(groups[key], cells)
	}

	return groups, nil
}
