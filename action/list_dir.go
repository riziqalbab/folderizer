package action

import (
	"context"
	"fmt"
	"folderizer/utils"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func ListDir(ctx context.Context, cmd *cli.Command) error {
	path_args := cmd.Args().Get(0)
	working_directory, _ := os.Getwd()

	if path_args == "" {
		file_list, err := utils.GetFileList(working_directory)

		if err != nil {
			log.Fatal(err)
		}

		for _, file := range file_list {
			fmt.Println(file.Name())
		}
	} else {
		file_list, err := utils.GetFileList(path_args)

		if err != nil {
			log.Fatal(err)
		}

		for _, file := range file_list {
			fmt.Println(file.Name())
		}
	}

	return nil
}
