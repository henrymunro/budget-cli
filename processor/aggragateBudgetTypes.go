package processor

// AggragatedEntry - Contains the aggragated information for a specific budget type for a month
type AggragatedEntry struct {
	MonthYear  string
	BudgetType string
	Count      int32
	Total      float32
	Average    float32
	Min        float32
	Max        float32
}

func aggragateBudgetTypes(entries []MappedEntry) []AggragatedEntry {

	var aggragatedEntries []AggragatedEntry

	for _, entry := range entries {
		index := findBudgetTypeIndex(entry.BudgetType, convertToMonthYear(entry.Date), aggragatedEntries)
		if index == -1 {
			aggragatedEntries = append(aggragatedEntries, newAggragatedEntry(entry))
		} else {
			applyEntryToAggragate(entry, &aggragatedEntries[index])
		}
	}

	return aggragatedEntries
}

func findBudgetTypeIndex(budgetType string, monthYear string, aggragatedEntries []AggragatedEntry) int {
	var index = -1

	for i, aggragatedEntry := range aggragatedEntries {
		if aggragatedEntry.BudgetType == budgetType && aggragatedEntry.MonthYear == monthYear {
			index = i
			break
		}
	}
	return index
}

func newAggragatedEntry(entry MappedEntry) AggragatedEntry {
	var newTypeEntry = AggragatedEntry{BudgetType: entry.BudgetType}
	newTypeEntry.MonthYear = convertToMonthYear(entry.Date)
	newTypeEntry.Count = 1
	newTypeEntry.Total = entry.Amount
	newTypeEntry.Average = entry.Amount
	newTypeEntry.Min = entry.Amount
	newTypeEntry.Max = entry.Amount
	return newTypeEntry
}

func applyEntryToAggragate(entry MappedEntry, aggragatedEntry *AggragatedEntry) {
	aggragatedEntry.Average = ((aggragatedEntry.Total * float32(aggragatedEntry.Count)) + entry.Amount) / (float32(aggragatedEntry.Count) + 1.0)
	aggragatedEntry.Total = aggragatedEntry.Total + entry.Amount
	aggragatedEntry.Count = aggragatedEntry.Count + 1
	if entry.Amount < aggragatedEntry.Min {
		aggragatedEntry.Min = entry.Amount
	}
	if entry.Amount > aggragatedEntry.Max {
		aggragatedEntry.Max = entry.Amount
	}
}
