package reader

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// NationwideParser csv reader for nationwide statements
type NationwideParser struct{}

type nationwideEntry struct {
	date            string
	transactionType string
	description     string
	paidOut         float32
	paidIn          float32
	balance         float32
}

func logger(msg string) {
	fmt.Println("Nationwide parser: ", msg)
}

func (n NationwideParser) parse(input string) Entries {
	logger("Parsing file as Nationwide csv")

	reader := csv.NewReader(strings.NewReader(input))
	reader.FieldsPerRecord = 6

	const nationwideHeaderLines = 3
	logger("Discarding first 3 lines")
	discardNumberOfCSVLines(nationwideHeaderLines, reader)

	columnTitles, err := reader.Read()
	check(err)
	logger("Header line - " + strings.Join(columnTitles, ", "))

	var entries Entries
	i := 0
	for {
		i++
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		check(err)
		var amountWithoutCurrency = trimLeftChars(line[3], 1)
		if amountWithoutCurrency != "" {
			amountFloat, err := strconv.ParseFloat(amountWithoutCurrency, 64)
			check(err)
			entries = append(entries, Entry{date: line[0], description: line[2], amount: float32(amountFloat)})
		}

	}
	logger("Parsed " + strconv.Itoa(i) + " lines")

	return entries
}
