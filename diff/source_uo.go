package diff

import (
	"errors"
	"math/big"
	"strings"
	"time"

	"github.com/extrame/xls"
	excel "github.com/zacwhy/go-diff-transactions/xls"
)

type uoSource struct {
}

func (uoSource uoSource) String() string {
	return "uo"
}

func (uoSource uoSource) parse(fileName string) (recordGroups, error) {
	workBook, err := xls.Open(fileName, "utf-8")

	if err != nil {
		return nil, err
	}

	sheet := workBook.GetSheet(0)

	if sheet == nil {
		return nil, errors.New("no worksheet")
	}

	rows := excel.ReadRows(sheet)

	return parseArray(rows)
}

func parseArray(rows [][]string) (map[string][][]string, error) {
	m := make(map[string][][]string)

	for i := 11; i < len(rows); i++ {
		cells := rows[i]

		date := cells[0]
		time, err := time.Parse("02 Jan 2006", date)

		if err != nil {
			return nil, err
		}

		amount := cells[6]
		dollars := new(big.Rat)
		dollars.SetString(amount)
		cents := dollars.Mul(dollars, big.NewRat(100, 1))
		descriptions := strings.Split(cells[2], "\n")

		key := time.Format("2006-01-02") + " " + cents.FloatString(0)

		if _, ok := m[key]; !ok {
			m[key] = [][]string{}
		}

		m[key] = append(m[key], []string{date, amount, descriptions[0]})
	}

	return m, nil
}
