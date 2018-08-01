package processor

import (
	"encoding/json"
	"io/ioutil"
	"sort"
)

type mapping struct {
	BudgetType string
	Mapping    string
}

type mappingsJSONFile struct {
	Mappings []mapping
}

// getMappings - reads in budget type mappings from the json file name provided
func getMappings(budgetTypeFilePath string) []mapping {

	// Read in file
	getMappingsLogger("reading in file " + budgetTypeFilePath)
	budgetTypeMappingsFile, err := ioutil.ReadFile(budgetTypeFilePath)
	check(err)

	// Parse as JSON
	var mappingsJSON mappingsJSONFile
	json.Unmarshal([]byte(budgetTypeMappingsFile), &mappingsJSON)
	mappings := mappingsJSON.Mappings

	return sortByMappingLength(mappings)
}

func sortByMappingLength(mappings []mapping) []mapping {

	sort.Slice(mappings, func(i, j int) bool {
		return len(mappings[i].Mapping) > len(mappings[j].Mapping)
	})

	return mappings
}
