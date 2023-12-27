/*
Create a CSV parser that takes a CSV string as input and converts it into a two-dimensional slice or map.
Use the strings package to split the input into rows and columns.
*/
package main

import (
	"encoding/csv"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("-----CSV Parser-----")
	var csvStr = `Id, Name,  Age
1, Chitra,  25
2, Nitin,   30
3, Rajashri,24
4, Shalini, 20`

	result, err := parseCSV(csvStr)
	if err != nil {
		fmt.Println("Error parsing the string", err)
		return
	}

	fmt.Println("Parsed CSV : ")
	for _, row := range result {
		fmt.Println(row)
	}
}
func parseCSV(csvstr string) ([][]string, error) {
	reader := csv.NewReader(strings.NewReader(csvstr))

	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Println("error reading the file ", err)
		return nil, err
	}
	return rows, nil
}
