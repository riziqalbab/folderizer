package utils

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v3"
)

func Organize(basePath string, ctx context.Context, cmd *cli.Command) error {

	folder_prefix := cmd.String("folder_prefix")
	folder_suffix := cmd.String("folder_suffix")

	file_list, _ := GetFileList(basePath)

	for _, file := range file_list {
		extension := strings.TrimPrefix(strings.ToLower(filepath.Ext(file.Name())), ".")

		if !file.Type().IsDir() {

			var category string = ""

			if folder_prefix != "" {
				category = folder_prefix + "_" + GetCategory(extension)
			} else {
				category = GetCategory(extension)
			}

			if folder_suffix != "" {
				category = category + "_" + folder_suffix
			}

			if category != "other" {
				targetDirectory := filepath.Join(basePath, category)

				targetPath := filepath.Join(targetDirectory, file.Name())
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
