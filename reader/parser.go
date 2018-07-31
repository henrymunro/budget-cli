package reader

// Parser interface to read in csv files
type csvParser interface {
	parse(string) string
}

func parseCSV(csvData string) Entries {

	parser := NationwideParser{}

	parsedData := parser.parse(csvData)

	return parsedData
}
