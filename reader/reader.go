package reader

import (
	"io/ioutil"
)

// ReadAndParseFile reads in csv file and parses the result
func ReadAndParseFile(fileName string) Entries {

	csvData, err := ioutil.ReadFile(fileName)
	check(err)

	parsedData := parseCSV(string(csvData))

	return parsedData

}
