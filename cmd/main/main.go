package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// Parse the flags
	var isInfoFlagSet bool
	flag.BoolVar(&isInfoFlagSet, "i", false, "Include additional info.")
	var isRecursiveFlagSet bool
	flag.BoolVar(&isRecursiveFlagSet, "r", false, "Recursively list all the sub-directories content.")

	flag.Parse()

	// Default behavior without any flags
	cwd, err := os.Getwd()
	if err != nil {
		panic("Couldn't get the current working directory.")
	}

	fmt.Println("CWD:", cwd)

	entries, err := os.ReadDir(cwd)
	if err != nil {
		panic("Couldn't read directory.")
	}

	var directories []string
	for _, entry := range entries {
		// if entry.IsDir() {
		directories = append(directories, entry.Name())
		// }
	}

	fmt.Println("")
	for _, dir := range directories {
		fmt.Println(dir)
	}
}
