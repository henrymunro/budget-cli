package reader

// Entry - single entry read in from statement file
type Entry struct {
	date        string
	amount      float32
	description string
}

// Entries - slice of entries read in from statement file
type Entries = []Entry
