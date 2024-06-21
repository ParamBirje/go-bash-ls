package main

import (
	"flag"
	"fmt"

	defaults "github.com/parambirje/go-bash-ls/internal/default"
)

func main() {

	// Parse the flags
	var isInfoFlagSet bool
	flag.BoolVar(&isInfoFlagSet, "i", false, "Include additional info.")
	var isRecursiveFlagSet bool
	flag.BoolVar(&isRecursiveFlagSet, "r", false, "Recursively list all the sub-directories content.")

	flag.Parse()

	// global directories list
	var directories []defaults.Directory

	if !isInfoFlagSet && !isRecursiveFlagSet {
		// Default behavior without any flags

		err := defaults.RunDefault(&directories)
		if err != nil {
			fmt.Println(err)
		}

	} else if isInfoFlagSet {
		// Additional info flag is set

		fmt.Println("Starting runner with additional info.")

	} else if isRecursiveFlagSet {
		// Recursive flag is set

		fmt.Println("Starting runner recursively.")
	}

	// Additionally handle when both flags are set
}
