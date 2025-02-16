package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Organize(basePath string) error {

	file_list, _ := GetFileList(basePath)

	for _, file := range file_list {
		extension := strings.TrimPrefix(strings.ToLower(filepath.Ext(file.Name())), ".")

		if !file.Type().IsDir() {
			category := GetCategory(extension)
			if category != "other" {
				targetDirectory := filepath.Join(basePath, category)

				targetPath := filepath.Join(targetDirectory, file.Name())

				fmt.Println(filepath.Join(basePath, file.Name()))

				if _, err := os.Stat(targetDirectory); os.IsNotExist(err) {
					CreateFolder(basePath, category)
				}

				err := os.Rename(filepath.Join(basePath, file.Name()), targetPath)

				if err != nil {
					return fmt.Errorf("error moving file: %w", err)
				}

				fmt.Printf("Moved %s to %s\n", file.Name(), targetPath)
			}

		}

	}

	return nil

}
