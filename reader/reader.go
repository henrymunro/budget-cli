package reader

import (
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// ReadAndParseFile reads in csv file and parses the result
func ReadAndParseFile(fileName string) {

	csvData, err := ioutil.ReadFile(fileName)
	check(err)

	parseCSV(string(csvData))

	// fmt.Println(parsedCSV)s
	// return "hi"

}
