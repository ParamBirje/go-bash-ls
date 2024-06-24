package recursive

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	defaults "github.com/parambirje/go-bash-ls/internal/default"
)

func RunRecursive(directories *[]defaults.Directory, cwd string, prefix string) error {
	// Recursive behavior with content from subdirectories

	entries, err := os.ReadDir(cwd)
	if err != nil {
		return errors.New("could not read directories")
	}

	for _, entry := range entries {

		newEntry := defaults.Directory{
			Path: filepath.Join(cwd, entry.Name()),
			Name: prefix + "/" + entry.Name(),
			Type: entry.Type().String(),
		}

		*directories = append(*directories, newEntry)

		if entry.IsDir() {
			err = RunRecursive(directories, newEntry.Path, prefix+"/"+entry.Name())

			if err != nil {
				fmt.Println("ERROR: Could not read directories in", newEntry.Path)
			}
		}
	}

	return nil
}
