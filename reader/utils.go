package reader

import (
	"encoding/csv"
	"fmt"
)

func discardNumberOfCSVLines(numberOfLines int, r *csv.Reader) {
	var discardedLines [][]string
	for i := 1; i <= numberOfLines; i++ {
		data, _ := r.Read()
		discardedLines = append(discardedLines, data)
	}
	fmt.Println(discardedLines)
}

func trimLeftChars(s string, n int) string {
	m := 0
	for i := range s {
		if m >= n {
			return s[i:]
		}
		m++
	}
	return s[:0]
}
