package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Organize(basePath string) error {

	return filepath.WalkDir(basePath, func(path string, directory os.DirEntry, err error) error {
		basePath = strings.TrimPrefix(basePath, "./")
		if err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Println(directory.Type())

		extension := strings.ToLower(filepath.Ext(directory.Name()))
		category := GetCategory(extension)

		targetDirectory := filepath.Join(basePath, category)

		targetPath := filepath.Join(targetDirectory, directory.Name())

		if _, err := os.Stat(targetDirectory); os.IsNotExist(err) {
			CreateFolder(basePath, category)
		}

		err = os.Rename(path, targetPath)

		if err != nil {
			return fmt.Errorf("error moving file: %w", err)
		}

		fmt.Printf("Moved %s to %s\n", directory.Name(), targetPath)

		return nil

	})
}
