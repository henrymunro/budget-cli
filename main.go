package main

import (
	"github.com/henrymunro/budget/processor"
	"github.com/henrymunro/budget/reader"
	"github.com/henrymunro/budget/writer"
)

const mappingsFilePath = "budgetTypeMappings.json"

func main() {
	entries := reader.ReadAndParseFile("statement.csv")
	mappedEntries, aggragatedBudgetTypes, otherGroup := processor.Process(entries, mappingsFilePath)

	writer.WriteOutputToCsv("output.csv", mappedEntries, aggragatedBudgetTypes, otherGroup)
}
