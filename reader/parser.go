package reader

// Parser interface to read in csv files
type csvParser interface {
	parse(string) string
}

// const dateFormat = "2006-01-02"

func parseCSV(csvData string) Entries {

	parser := NationwideParser{}

	parsedData := parser.parse(csvData)

	return parsedData
}
