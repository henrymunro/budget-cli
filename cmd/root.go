package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/henrymunro/budget/processor"
	"github.com/henrymunro/budget/reader"
	"github.com/henrymunro/budget/writer"
	"github.com/spf13/cobra"
)

const defaultOutputFile = "output.csv"
const defaultMappingsFilePath = "budgetTypeMappings.json"

var inputFile string
var outputFile string
var mappingsFilePath string

var rootCmd = &cobra.Command{
	Use:   "budget-cli",
	Short: "Processes a statment, maps descriptions to budget types and aggragates per month",
	Long:  "Produces a dependency graph of services from a package.json",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("You must pass the path to the statement file to process")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		inputFile := args[0]

		if len(outputFile) == 0 {
			outputFile = defaultOutputFile
		}
		if len(mappingsFilePath) == 0 {
			mappingsFilePath = defaultMappingsFilePath
		}

		entries := reader.ReadAndParseFile(inputFile)
		mappedEntries, aggragatedBudgetTypes, otherGroup := processor.Process(entries, mappingsFilePath)

		writer.WriteOutputToCsv(outputFile, mappedEntries, aggragatedBudgetTypes, otherGroup)
	},
}

func init() {
	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Path to the output file - default output.csv")
	rootCmd.Flags().StringVarP(&mappingsFilePath, "mappings", "m", "", "Optional path to mappings file - default budgetTypeMappings.json")
}

// Execute enters the command line interface
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
