package main

import (
	"flag"
	"fmt"

	defaults "github.com/parambirje/go-bash-ls/internal/default"
	"github.com/parambirje/go-bash-ls/internal/recursive"
)

func main() {

	// Parse the flags
	var isInfoFlagSet bool
	flag.BoolVar(&isInfoFlagSet, "i", false, "Include additional info.")
	var isRecursiveFlagSet bool
	flag.BoolVar(&isRecursiveFlagSet, "r", false, "Recursively list all the sub-directories content.")

	flag.Parse()

	// Global directories list
	var directories []defaults.Directory

	// Getting current working directory
	cwd, err := defaults.GetCurrentDirectory()
	if err != nil {
		fmt.Println(err)
		return
	}

	if isRecursiveFlagSet {
		// Recursive flag is set
		err := recursive.RunRecursive(&directories, cwd, ".")
		if err != nil {
			fmt.Println(err)
		}

	} else {
		// Default behavior without any flags
		err := defaults.RunDefault(&directories, cwd)
		if err != nil {
			fmt.Println(err)
		}
	}

	// Printing the directories list
	fmt.Println()
	for _, dir := range directories {
		if isRecursiveFlagSet && string(dir.Type[0]) == "d" {
			// for spacing between directories
			fmt.Println()
		}

		if isInfoFlagSet {
			fmt.Println(dir.Type, dir.Name)
		} else {
			fmt.Println(dir.Name)
		}
	}
}
