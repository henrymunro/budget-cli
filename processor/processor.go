package processor

import (
	"github.com/henrymunro/budget/reader"
)

type entries = reader.Entries

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Process - applies the budget type mappings and aggragates on a monthly basis
func Process(entries entries, budgetTypeFilePath string) ([]MappedEntry, []AggragatedEntry) {

	mappings := getMappings(budgetTypeFilePath)
	mappedEnties := applyMappings(entries, mappings)
	aggragatedEntries := aggragateBudgetTypes(mappedEnties)

	return mappedEnties, aggragatedEntries
}
