package writer

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/henrymunro/budget/processor"
)

type mappedEntry = processor.MappedEntry
type aggragatedBudgetType = processor.AggragatedBudgetType

type csvRow = []string

const dateFormatLong = "2006-01-02 15:04"
const dateFormatShort = "2006-01-02"

// WriteOutputToCsv - write the processed mappedEntries and aggragatedBudgetTypes to specified output file
func WriteOutputToCsv(filepath string, mappedEntries []mappedEntry, aggragatedBudgetTypes []aggragatedBudgetType) {
	logger("Writing to file " + filepath)
	file, err := os.Create(filepath)
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	var now = time.Now()
	writeCommentRow("Proccessed at "+now.Format(dateFormatLong), writer)
	writeCommentRow("", writer)

	logger("Writing aggragated budget types")
	writeAggragatedBudgetTypes(aggragatedBudgetTypes, writer)

	writeCommentRow("", writer)
	logger("Writing mapped entries")
	writeMappedEntries(mappedEntries, writer)

}

func writeAggragatedBudgetTypes(aggragatedBudgetTypes []aggragatedBudgetType, writer *csv.Writer) {
	var header = csvRow{"Date", "BudgetType", "Total", "Count", "Average", "Min", "Max"}
	writeRow(header, writer)

	for _, value := range aggragatedBudgetTypes {
		var row = make([]string, 7)
		row[0] = value.MonthYear
		row[1] = value.BudgetType
		row[2] = "£" + floatToString(value.Total)
		row[3] = strconv.Itoa(int(value.Count))
		row[4] = "£" + floatToString(value.Average)
		row[5] = "£" + floatToString(value.Min)
		row[6] = "£" + floatToString(value.Max)
		writeRow(row, writer)
	}
}

func writeMappedEntries(mappedEntires []mappedEntry, writer *csv.Writer) {
	var header = csvRow{"Date", "BudgetType", "Amount", "Description", "Mapping"}
	writeRow(header, writer)

	for _, value := range mappedEntires {
		var row = make([]string, 5)
		row[0] = value.Date.Format(dateFormatShort)
		row[1] = value.BudgetType
		row[2] = "£" + floatToString(value.Amount)
		row[3] = value.Description
		row[4] = value.Mapping
		writeRow(row, writer)
	}
}

func floatToString(num float32) string {
	s32 := strconv.FormatFloat(float64(num), 'f', 2, 32)
	return s32
}

func writeCommentRow(comment string, writer *csv.Writer) {
	var row = make([]string, 1)
	row[0] = comment
	writeRow(row, writer)
}

func writeRow(row []string, writer *csv.Writer) {
	err := writer.Write(row)
	checkError("Cannot write to file", err)
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func logger(msg string) {
	fmt.Println("Writer: ", msg)
}
