package main

import (
	"github.com/henrymunro/budget/processor"
	"github.com/henrymunro/budget/reader"
)

const mappingsFilePath = "budgetTypeMappings.json"

func main() {
	entries := reader.ReadAndParseFile("statement.csv")
	mappedEntries, aggragatedBudgetTypes := processor.Process(entries, mappingsFilePath)
}
