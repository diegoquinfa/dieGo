package table

import (
	"fmt"
	"strings"
)

const (
	ESQUINA_INF_DER = "┘"
	ESQUINA_INF_IZQ = "└"

	ESQUINA_SUP_DER = "┐"
	ESQUINA_SUP_IZQ = "┌"

	PARED_HORIZONTAL = "─"
	PARED_VERTICAL   = "│"

	INTERSECCION_DER = "┤"
	INTERSECCION_IZQ = "├"

	INTERSECCION_SUP = "┬"
	INTERSECCION_MED = "┼"
	INTERSECCION_INF = "┴"
)

type Row struct {
	Columns      []any
	maxLenColumn []int
}

func GetMaxWidthColumns(rows []Row) []int {
	maxWidths := make([]int, len(rows[0].maxLenColumn))

	for _, row := range rows {
		for j, len := range row.maxLenColumn {
			if maxWidths[j] < len {
				maxWidths[j] = len
			}
		}
	}

	return maxWidths
}

func CreateTable(rows []Row) {
	if len(rows) == 0 {
		return
	}

	maxWidths := GetMaxWidthColumns(rows)

	fmt.Print(ESQUINA_SUP_IZQ)

	for i, width := range maxWidths {
		for i := 0; i < width+2; i++ {
			fmt.Print(PARED_HORIZONTAL)
		}
		if i != len(maxWidths)-1 {
			fmt.Print(INTERSECCION_SUP)
		}
	}

	fmt.Println(ESQUINA_SUP_DER)

	for i, row := range rows {
		if i == 0 {
			for j, column := range row.Columns {
				text := fmt.Sprintf("%v", column)
				len := maxWidths[j] - row.maxLenColumn[j]
				fmt.Printf("%s %s", PARED_VERTICAL, text)
				for k := 0; k < len+1; k++ {
					fmt.Print(" ")
				}
			}
			fmt.Println(PARED_VERTICAL)

			fmt.Print(INTERSECCION_IZQ)
			for j, width := range maxWidths {
				for k := 0; k < width+2; k++ {
					fmt.Print(PARED_HORIZONTAL)
				}
				if j != len(maxWidths)-1 {
					fmt.Print(INTERSECCION_MED)
				}
			}
			fmt.Println(INTERSECCION_DER)
			continue
		}

		for j, column := range row.Columns {
			text := fmt.Sprintf("%v", column)
			spaces := maxWidths[j] - row.maxLenColumn[j]
			fmt.Printf("%s %s", PARED_VERTICAL, text)
			for k := 0; k < spaces+1; k++ {
				fmt.Print(" ")
			}
		}
		fmt.Println(PARED_VERTICAL)

		if i != len(rows)-1 {
			for j := range row.Columns {
				fmt.Printf("%s", PARED_VERTICAL)
				spaces := maxWidths[j]
				for k := 0; k < spaces+2; k++ {
					fmt.Print(" ")
				}
			}

			fmt.Println(PARED_VERTICAL)
		}
	}

	fmt.Print(ESQUINA_INF_IZQ)
	for i, width := range maxWidths {
		for i := 0; i < width+2; i++ {
			fmt.Print(PARED_HORIZONTAL)
		}
		if i != len(maxWidths)-1 {
			fmt.Print(INTERSECCION_INF)
		}
	}
	println(ESQUINA_INF_DER)
}

func NewRow(columns ...any) *Row {
	row := Row{}

	row.Columns = append(row.Columns, columns...)

	for i := range row.Columns {
		text := fmt.Sprintf("%v", columns[i])
		lenText := len([]rune(text))
		if strings.Contains(text, "\x1b[31m") || strings.Contains(text, "\x1b[32m") {
			lenText -= 9
		}
		row.maxLenColumn = append(row.maxLenColumn, lenText)
	}

	return &row
}
