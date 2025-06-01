package diff

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strings"

	"github.com/extrame/xls"
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

			if text == `"Transaction Date/Time","Amount","Merchant Name"` {
				return ezSource{}, nil
			}

			re := regexp.MustCompile(`^[0-9]{2}/[0-9]{2}/[0-9]{1,},.+ [0-9]{2}/[0-9]{2}/[0-9]{1,} •••• •••• •••• [0-9]{4} -[0-9]{1,}\.[0-9]{2} SGD,-[0-9]{1,}\.[0-9]{2}`)
			if re.MatchString(text) {
				return hsSource{}, nil
			}

			headers := strings.Split(text, ",")
			dateIndex := array.IndexOf("d", headers)
			amountIndex := array.IndexOf("a", headers)
			descriptionIndex := array.IndexOf("c", headers)
			if dateIndex != -1 && amountIndex != -1 && descriptionIndex != -1 {
				return localSource{}, nil
			}
		}

		if i == 4 {
			if text == "Date,DESCRIPTION,Foreign Currency Amount,SGD Amount" {
				return scSource{}, nil
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if source, err := tryXlsFile(fileName); err == nil {
		return source, nil
	}

	return nil, errors.New("cannot detect source type for: " + fileName)
}

func tryXlsFile(fileName string) (source, error) {
	workBook, err := xls.Open(fileName, "utf-8")

	if err != nil {
		return nil, err
	}

	return tryUoXlsFile(workBook)
}

func tryUoXlsFile(workBook *xls.WorkBook) (source, error) {
	sheet := workBook.GetSheet(0)

	if sheet == nil {
		return nil, errors.New("no worksheet")
	}

	if sheet.MaxRow < 9 {
		return nil, errors.New("no headers")
	}

	row := sheet.Row(9)

	isUo := row != nil &&
		row.Col(0) == "Transaction Date" &&
		row.Col(1) == "Posting Date" &&
		row.Col(2) == "Description" &&
		row.Col(3) == "Foreign Currency Type" &&
		row.Col(4) == "Transaction Amount(Foreign)" &&
		row.Col(5) == "Local Currency Type" &&
		row.Col(6) == "Transaction Amount(Local)"

	if !isUo {
		return nil, errors.New("headers do not match")
	}

	return uoSource{}, nil
}
