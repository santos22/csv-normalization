package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var columns []string

func main() {
	ignoreTitle := true
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		row := scanner.Text()
		if ignoreTitle {
			ignoreTitle = false
			continue
		}

		// count of commas before we reach Timestamp,Address
		commaCount := 0
		// returns Timestamp,Address
		example := retrieveColumnValues(commaCount, row)

		// returns index of ',' between Timestamp and Address
		i := strings.Index(example, ",")

		// convert string to rune slice
		runes := []rune(example)

		printColumns(columns)
		fmt.Println(string(runes[i+1:])) // get Address value
		fmt.Println(string(runes[:i]))   // get Timestamp value
		fmt.Println("=============================================")
		columns = columns[:0] // clear columns slice
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func printColumns(columns []string) {
	for _, value := range columns {
		fmt.Println(value)
	}
}

// start from right of row to retrieve column values
func retrieveColumnValues(commaCount int, row string) string {
	if commaCount == 6 {
		return row
	}

	i := strings.LastIndex(row, ",")
	columnValue := row[i+1 : len(row)]

	columns = append(columns, columnValue)
	commaCount = commaCount + 1
	return retrieveColumnValues(commaCount, row[:i])
}
