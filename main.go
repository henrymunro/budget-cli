package main

import (
	"github.com/henrymunro/budget/reader"
)

func main() {
	reader.ReadAndParseFile("statement.csv")
}
