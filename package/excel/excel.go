package excel

import (
	"errors"

	"github.com/xuri/excelize/v2"
)

var (
	colLetters = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
)

// getSheetName returns the name of the sheet
func getSheetName(ex *excelize.File) (string, error) {
	sheets := ex.GetSheetList()
	if len(sheets) > 1 {
		return "", errors.New("the excel file cannot have more one sheet")
	} else {
		return sheets[0], nil
	}
}

// NextCol function returns the next index and next excel column name
// It receives the current index and number of steps
func NextCol(currentIndex, steps int) (int, string) {
	l := len(colLetters)
	//fmt.Println(l)
	nextIndex := currentIndex + steps
	nextColName := ""
	if nextIndex >= l {

		p1 := nextIndex/l - 1
		p2 := nextIndex % l
		nextColName = colLetters[p1] + colLetters[p2]
		//fmt.Printf("current index = %d, next index = %d, col name = %s\n", currentIndex, nextIndex, nextColName)

		return nextIndex, nextColName
	}
	nextColName = colLetters[nextIndex]
	//fmt.Printf("current index = %d, next index = %d, col name = %s\n", currentIndex, nextIndex, nextColName)

	return nextIndex, nextColName
}
