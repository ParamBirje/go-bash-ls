package defaults

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type Directory struct {
	Path string
	Type string // "f" or "d" for file or directory
}

func RunDefault(directories *[]Directory) error {
	// Default behavior without any flags
	cwd, err := os.Getwd()
	if err != nil {
		return errors.New("could not get the current working directory")
	}

	fmt.Println("CWD:", cwd)

	entries, err := os.ReadDir(cwd)
	if err != nil {
		return errors.New("could not read directories")
	}

	// var directories []Directory
	for _, entry := range entries {
		*directories = append(*directories, Directory{
			Path: filepath.Join(cwd, entry.Name()),
			Type: entry.Type().String(),
		})
	}

	fmt.Println("")
	for _, dir := range *directories {
		fmt.Println(dir)
	}

	return nil
}
