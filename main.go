package main

import (
	"fmt"
	"os"
)

func main() {
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

	fmt.Println("\nChildren:")
	for _, dir := range directories {
		fmt.Println(dir)
	}
}
