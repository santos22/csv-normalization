package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
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

		//printColumns(columns)
		columns = append(columns, string(runes[i+1:]))
		columns = append(columns, string(runes[:i]))
		printColumns(columns)
		//fmt.Println(string(runes[i+1:])) // get Address value
		//fmt.Println(string(runes[:i]))   // get Timestamp value
		fmt.Println("=============================================")
		columns = columns[:0] // clear columns slice
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func printColumns(columns []string) {
	// Notes 0
	// TotalDuration 1
	// BarDuration 2
	// FooDuration 3
	// FullName 4
	// ZIP 5
	// Address 6
	// Timestamp 7

	// for _, value := range columns {
	// 	fmt.Println(value)
	// }
	//convertToEasternTime(columns[7])
	//prefixZipcode(columns[5])
	//fmt.Println(strings.ToUpper(columns[4]))
	fmt.Println(columns[2])
	//t, _ := time.Parse("5:04:05.000Z", columns[2])
	//fmt.Println(t.Format("150405"))
	//fmt.Println(t)
	duration := strings.Split(columns[2], ":")
	fmt.Println(duration)
	hours := duration[0]
	minutes := duration[1]
	seconds := strings.Split(duration[2], ".")[0]
	microseconds := strings.Split(duration[2], ".")[1]
	fmt.Println(hours + minutes + seconds + microseconds)
}

func convertToSeconds(durationTime string) {

}

func prefixZipcode(zipcode string) {
	var buffer bytes.Buffer

	for i := 0; i < 5-len(zipcode); i++ {
		buffer.WriteString("0")
	}
	buffer.WriteString(zipcode)

	fmt.Println(buffer.String())
}

func convertToEasternTime(timeStamp string) {
	if len(timeStamp) == 18 {
		t, _ := time.Parse("1/02/06 3:04:05 PM", timeStamp)
		if t.Format(time.RFC3339) == "0001-01-01T00:00:00Z" {
			t, _ := time.Parse("01/2/06 03:04:05 PM", timeStamp)
			if t.Format(time.RFC3339) == "0001-01-01T00:00:00Z" {
				t, _ := time.Parse("1/2/06 03:04:05 PM", timeStamp)
				if t.Format(time.RFC3339) == "0001-01-01T00:00:00Z" {
					t, _ := time.Parse("01/2/06 3:04:05 PM", timeStamp)
					easternTime := t.Add(time.Hour * 3)
					fmt.Println(easternTime.Format(time.RFC3339))
				} else {
					easternTime := t.Add(time.Hour * 3)
					fmt.Println(easternTime.Format(time.RFC3339))
				}
			} else {
				easternTime := t.Add(time.Hour * 3)
				fmt.Println(easternTime.Format(time.RFC3339))
			}
		} else {
			easternTime := t.Add(time.Hour * 3)
			fmt.Println(easternTime.Format(time.RFC3339))
		}
	} else if len(timeStamp) == 19 {
		t, _ := time.Parse("1/02/06 03:04:05 PM", timeStamp)
		if t.Format(time.RFC3339) == "0001-01-01T00:00:00Z" {
			t, _ := time.Parse("01/02/06 3:04:05 PM", timeStamp)
			if t.Format(time.RFC3339) == "0001-01-01T00:00:00Z" {
				t, _ := time.Parse("01/2/06 03:04:05 PM", timeStamp)
				easternTime := t.Add(time.Hour * 3)
				fmt.Println(easternTime.Format(time.RFC3339))
			} else {
				easternTime := t.Add(time.Hour * 3)
				fmt.Println(easternTime.Format(time.RFC3339))
			}
		} else {
			easternTime := t.Add(time.Hour * 3)
			fmt.Println(easternTime.Format(time.RFC3339))
		}
	} else if len(timeStamp) == 20 {
		t, _ := time.Parse("01/02/06 03:04:05 PM", timeStamp)
		easternTime := t.Add(time.Hour * 3)
		fmt.Println(easternTime.Format(time.RFC3339))
	}
}

// start from right of row to retrieve column values
func retrieveColumnValues(commaCount int, column string) string {
	if commaCount == 6 {
		return column
	}

	// get string between quotes
	r, _ := regexp.Compile("\"([^\"]*)\"")

	var columnValue string
	//var i int
	indices := r.FindStringIndex(column)
	if indices != nil {
		if indices[1] == len(column) {
			// Notes
			fmt.Println(indices)
			fmt.Println(len(column))
			fmt.Println(column[indices[0]:indices[1]])
			fmt.Println(indices[0] - 1)
			columnValue = column[indices[0]:indices[1]]
			//i = indices[0] - 1
		}
	}

	i := strings.LastIndex(column, ",")
	//fmt.Println(i)
	columnValue = column[i+1 : len(column)]

	columns = append(columns, columnValue)
	commaCount = commaCount + 1
	return retrieveColumnValues(commaCount, column[:i])
}
