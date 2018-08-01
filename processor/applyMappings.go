package processor

import (
	"regexp"
	"strings"
	"time"
)

// MappedEntry - a statement row with the budget type mapping applied
type MappedEntry struct {
	Description string
	Date        time.Time
	Amount      float32
	BudgetType  string
	Mapping     string
}

const other = "Other"

func applyMappings(entries entries, mappings []mapping) []MappedEntry {
	applyMappingsLogger("Applying mappings")

	var mappedEntries []MappedEntry
	for _, entry := range entries {
		budgetType, mapping := getMappingForDescription(entry.Description, mappings)
		mappedEntries = append(mappedEntries, MappedEntry{Description: entry.Description, Date: entry.Date, Amount: entry.Amount, BudgetType: budgetType, Mapping: mapping})
	}

	return mappedEntries
}

func getMappingForDescription(description string, mappings []mapping) (string, string) {
	var budgetType = other
	var matchedMapping = ""
	for _, mapping := range mappings {
		uppercaseMapping := strings.ToUpper(mapping.Mapping)
		uppercaseDescription := strings.ToUpper(description)
		match, _ := regexp.MatchString(uppercaseMapping, uppercaseDescription)
		if match {
			budgetType = mapping.BudgetType
			matchedMapping = mapping.Mapping
			break
		}
	}
	return budgetType, matchedMapping
}
