package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/arnemolland/exinc/interval"
)

// Intervals struct to hold both includes and excludes
type Intervals struct {
	Includes []interval.Interval `json:"includes"`
	Excludes []interval.Interval `json:"excludes"`
}

var (
	inputFile string
)

func main() {
	cmd := &cobra.Command{Use: "exinc", Short: "Takes a json input and merges the intervals", Run: func(cmd *cobra.Command, args []string) {
		// Create a reader based on stdin or file input
		var decoder *json.Decoder
		if inputFile != "" {
			file, err := os.Open(inputFile)
			if err != nil {
				fmt.Println("Error opening file:", err)
				os.Exit(1)
			}
			defer file.Close()
			decoder = json.NewDecoder(file)
		} else {
			decoder = json.NewDecoder(os.Stdin)
		}

		var intervals Intervals
		if err := decoder.Decode(&intervals); err != nil {
			fmt.Println("Error parsing JSON:", err)
			os.Exit(1)
		}

		result := interval.ProcessIntervals(intervals.Includes, intervals.Excludes)

		for _, interval := range result {
			fmt.Printf("%d-%d\n", interval.Start, interval.End)
		}
	}}

	cmd.Flags().StringVarP(&inputFile, "file", "f", "", "input file containing json")

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
