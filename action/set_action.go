package action

import (
	"context"
	"errors"
	"fmt"
	"folderizer/utils"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v3"
)

func SetAction(ctx context.Context, cmd *cli.Command) error {
	pathsArg := cmd.Args().Get(0)
	workingDirectoryPath, _ := os.Getwd()

	file_prefix_name := cmd.String("file_prefix")
	file_suffix_name := cmd.String("file_suffix")
	folder_destination := cmd.String("to_folder")

	if file_prefix_name == "" && file_suffix_name == "" {
		err := errors.New("please specify a file prefix or a file suffix. For example: folderizer set --file_prefix <file_prefix_name> --file_suffix <file_suffix_name> --to_folder <folder_destination>")
		return fmt.Errorf("error occurred: %v", err)
	}
	if folder_destination == "" {
		err := errors.New("please specify a folder destination. For example: folderizer set --file_prefix <file_prefix_name> --file_suffix <file_suffix_name> --to_folder <folder_destination>")
		return fmt.Errorf("error occurred: %v", err)

	}

	if pathsArg != "" {
		fmt.Printf("Organizing %s\n", pathsArg)
		organizeSet(pathsArg, cmd)
	} else {
		fmt.Printf("Organizing %s\n", workingDirectoryPath)
		organizeSet(workingDirectoryPath, cmd)
	}

	return nil
}

func organizeSet(basePath string, cmd *cli.Command) {

	var file_suffix_name string = cmd.String("file_suffix")
	var file_prefix_name string = cmd.String("file_prefix")

	files_list, _ := utils.GetFileList(basePath)
	// fmt.Println(file_prefix_name, file_suffix_name, folder_destination)
	for _, file := range files_list {

		if !file.Type().IsDir() {

			extension := filepath.Ext(file.Name())

			file_name := strings.TrimSuffix(file.Name(), extension)

			if (file_prefix_name != "" && strings.HasPrefix(file_name, file_prefix_name)) ||
				(file_suffix_name != "" && strings.HasSuffix(file_name, file_suffix_name)) {

				targetPath := filepath.Join(basePath, cmd.String("to_folder"), file.Name())
				if _, err := os.Stat(filepath.Join(basePath, cmd.String("to_folder"))); os.IsNotExist(err) {
					os.Mkdir(filepath.Join(basePath, cmd.String("to_folder")), 0777)
				}

				os.Rename(filepath.Join(basePath, file.Name()), targetPath)
				fmt.Println("File " + file.Name() + " has been moved to " + cmd.String("to_folder"))

			}
		}

	}

}
