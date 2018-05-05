package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var rows []string

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

		// convert string to rune slice to retrieve substrings
		runes := []rune(example)

		rows = append(rows, string(runes[i+1:])) // append Address value
		rows = append(rows, string(runes[:i]))   // append Timestamp value
		printCSV(rows)
		fmt.Println("=============================================")
		rows = rows[:0] // clear rows slice
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func printCSV(rows []string) {
	convertToEasternTime("Timestamp: " + rows[7])        // Timestamp 7
	fmt.Println("Address: " + rows[6])                   // Address 6
	prefixZipcode("ZIP: " + rows[5])                     // ZIP 5
	fmt.Println("FullName: " + strings.ToUpper(rows[4])) // FullName 4

	// FooDuration 3
	// fmt.Println(rows[3])
	// duration := strings.Split(rows[3], ":")
	// fmt.Println(duration)
	// hours := duration[0]
	// minutes := duration[1]
	// seconds := strings.Split(duration[3], ".")[0]
	// microseconds := strings.Split(duration[3], ".")[1]
	// convertToSeconds(hours, minutes, seconds, microseconds)

	// BarDuration 2 in HH:MM:SS.MS format
	duration := strings.Split(rows[2], ":")
	hours := duration[0]
	minutes := duration[1]
	seconds := strings.Split(duration[2], ".")[0]
	microseconds := strings.Split(duration[2], ".")[1]
	convertToSeconds(hours, minutes, seconds, microseconds)

	fmt.Println("TotalDuration: " + rows[1]) // TotalDuration 1
	fmt.Println("Notes: " + rows[0])         // Notes 0
}

// Convert from HH:MM:SS.MS format to floating point seconds format
func convertToSeconds(hours, minutes, seconds, microseconds string) {
	hoursInSeconds, _ := strconv.Atoi(hours)
	minutesInSeconds, _ := strconv.Atoi(minutes)
	formattedSeconds, _ := strconv.Atoi(seconds)
	formattedSeconds = formattedSeconds + (hoursInSeconds * 3600) + (minutesInSeconds * 60)

	var buffer bytes.Buffer

	buffer.WriteString(strconv.Itoa(formattedSeconds))
	buffer.WriteString(".")
	buffer.WriteString(microseconds)

	fmt.Println("BarDuration: " + buffer.String())
}

// Format ZIP codes as 5 digits
// If there are less than 5 digits, assume 0 as the prefix
func prefixZipcode(zipcode string) {
	var buffer bytes.Buffer

	for i := 0; i < 5-len(zipcode); i++ {
		buffer.WriteString("0")
	}
	buffer.WriteString(zipcode)

	fmt.Println(buffer.String())
}

// Format timestamp values to ISO-8601
// and convert to US/Eastern
// Accounts for 8 different formats
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

// Start from right of row to recursively retrieve column values
// TODO: Handle retrieving text between quotes
func retrieveColumnValues(commaCount int, column string) string {
	if commaCount == 6 {
		return column
	}

	// get string between quotes
	r, _ := regexp.Compile("\"([^\"]*)\"")

	var columnValue string
	indices := r.FindStringIndex(column)
	if indices != nil {
		// If end index equal to end of row, Notes field is in quotes
		if indices[1] == len(column) {
			fmt.Println(indices)
			fmt.Println(len(column))
			fmt.Println(column[indices[0]:indices[1]])
			fmt.Println(indices[0] - 1)
			columnValue = column[indices[0]:indices[1]]
		}
	}

	i := strings.LastIndex(column, ",")
	columnValue = column[i+1 : len(column)]

	rows = append(rows, columnValue)
	commaCount = commaCount + 1
	return retrieveColumnValues(commaCount, column[:i])
}
