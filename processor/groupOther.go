package processor

import "sort"

// OtherGroup - gives a count for a description that could not be mapped
type OtherGroup struct {
	Description string
	Count       int
}

func groupOther(mappedEntries []MappedEntry) []OtherGroup {
	descriptionsToGroup := filterForOtherBudgetTypes(mappedEntries)

	var aggragatedOtherGroups []OtherGroup
	for _, description := range descriptionsToGroup {
		index := findOtherGroupIndex(description, aggragatedOtherGroups)
		if index == -1 {
			newOtherGroup := OtherGroup{Description: description, Count: 1}
			aggragatedOtherGroups = append(aggragatedOtherGroups, newOtherGroup)
		} else {
			aggragatedOtherGroups[index].Count = aggragatedOtherGroups[index].Count + 1
		}
	}

	sortedOtherGroups := sortOtherGroupsByCount(aggragatedOtherGroups)

	return sortedOtherGroups
}

func filterForOtherBudgetTypes(mappedEntries []MappedEntry) []string {
	var output []string
	for _, entry := range mappedEntries {
		if entry.BudgetType == other {
			output = append(output, entry.Description)
		}
	}
	return output
}

func findOtherGroupIndex(description string, otherGroups []OtherGroup) int {
	var index = -1
	for i, otherGroup := range otherGroups {
		if otherGroup.Description == description {
			index = i
			break
		}
	}
	return index
}

func sortOtherGroupsByCount(otherGroups []OtherGroup) []OtherGroup {
	sort.Slice(otherGroups, func(i, j int) bool {
		if otherGroups[i].Count != otherGroups[j].Count {
			return otherGroups[i].Count > otherGroups[j].Count
		}

		return otherGroups[i].Description < otherGroups[j].Description
	})

	return otherGroups
}
