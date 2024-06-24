package defaults

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type Directory struct {
	Path string
	Name string
	Type string // "f" or "d" for file or directory
}

func RunDefault(directories *[]Directory, cwd string) error {
	// Default behavior without any flags
	// Directly mutates the directories list using pointer

	entries, err := os.ReadDir(cwd)
	if err != nil {
		return errors.New("could not read directories")
	}

	// var directories []Directory
	for _, entry := range entries {
		*directories = append(*directories, Directory{
			Path: filepath.Join(cwd, entry.Name()),
			Name: entry.Name(),
			Type: entry.Type().String(),
		})
	}

	return nil
}

func GetCurrentDirectory() (path string, err error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", errors.New("could not get the current working directory")
	}

	fmt.Println("CWD:", cwd)
	return cwd, nil
}
