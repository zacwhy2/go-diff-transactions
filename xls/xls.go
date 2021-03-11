package xls

import "github.com/extrame/xls"

// ReadRows reads all rows in a worksheet
func ReadRows(sheet *xls.WorkSheet) [][]string {
	rows := make([][]string, sheet.MaxRow+1)

	for i := 0; i <= int(sheet.MaxRow); i++ {
		row := sheet.Row(i)

		rows[i] = make([]string, row.LastCol()+1)

		for j := row.FirstCol(); j <= row.LastCol(); j++ {
			rows[i][j] = row.Col(j)
		}
	}

	return rows
}
