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
	xlFile, err := xls.Open(fileName, "utf-8")

	if err != nil {
		return nil, err
	}

	sheet1 := xlFile.GetSheet(0)

	if sheet1 == nil {
		return nil, errors.New("no worksheet")
	}

	rows := excel.ReadRows(sheet1)

	return parseArray(rows)
}

func parseArray(rows [][]string) (map[string][][]string, error) {
	m := make(map[string][][]string)

	for i := 11; i < len(rows); i++ {
		row := rows[i]

		date := row[0]
		time, err := time.Parse("02 Jan 2006", date)

		if err != nil {
			return nil, err
		}

		amount := row[6]
		dollars := new(big.Rat)
		dollars.SetString(amount)
		cents := dollars.Mul(dollars, big.NewRat(100, 1))

		key := time.Format("2006-01-02") + " " + cents.FloatString(0)

		description := row[2]
		descriptions := strings.Split(description, "\n")

		var arr []string
		arr = append(arr, date, amount, descriptions[0])

		if _, ok := m[key]; !ok {
			m[key] = [][]string{}
		}

		m[key] = append(m[key], arr)
	}

	return m, nil
}
