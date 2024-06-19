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

	fmt.Println(cwd)
}
