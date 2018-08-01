package reader

import (
	"encoding/csv"
	"io"
	"strconv"
	"strings"
	"time"
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

func (n NationwideParser) parse(input string) Entries {
	nationwideLogger("Parsing file as Nationwide csv")

	reader := csv.NewReader(strings.NewReader(input))
	reader.FieldsPerRecord = 6

	const nationwideHeaderLines = 3
	nationwideLogger("Discarding first 3 lines")
	discardNumberOfCSVLines(nationwideHeaderLines, reader)

	columnTitles, err := reader.Read()
	check(err)
	nationwideLogger("Header line - " + strings.Join(columnTitles, ", "))

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
			entries = append(entries, Entry{Date: convertDatesToShortFormat(line[0]), Description: line[2], Amount: float32(amountFloat)})
		}

	}
	nationwideLogger("Parsed " + strconv.Itoa(i) + " lines")

	return entries
}

func convertDatesToShortFormat(date string) time.Time {
	const shortForm = "02 Jan 2006"
	t, err := time.Parse(shortForm, date)
	check(err)
	return t
}
