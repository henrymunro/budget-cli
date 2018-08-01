package reader

import "time"

// Entry - single entry read in from statement file
type Entry struct {
	Date        time.Time
	Amount      float32
	Description string
}

// Entries - slice of entries read in from statement file
type Entries = []Entry
