package processor

import (
	"fmt"
	"time"
)

func rootLogger(msg string) {
	fmt.Println("Processor: ", msg)
}

func getMappingsLogger(msg string) {
	rootLogger("Get mappings: " + msg)
}

func applyMappingsLogger(msg string) {
	rootLogger("Apply mappings: " + msg)
}

func convertToMonthYear(date time.Time) string {
	monthYear := date.Format("2006-01")
	return monthYear
}
