package reader

type entry struct {
	date        string
	amount      float32
	description string
}

// Parser interface to read in csv files
type csvParser interface {
	parse(string) string
}

func parseCSV(csvData string) []entry {

	parser := NationwideParser{}

	parsedData := parser.parse(csvData)

	// for _, row := range parsedData {
	// 	fmt.Println(row)
	// }

	return parsedData
}
