package action

import (
	"context"
	"folderizer/utils"
	"os"

	"github.com/urfave/cli/v3"
)

func Test(ctx context.Context, cmd *cli.Command) error {
	path_args := cmd.Args().Get(0)
	working_directory, _ := os.Getwd()

	if path_args == "" {
		utils.CreateFolder(working_directory, "test")
	} else {
		utils.CreateFolder(path_args, "test")
	}

	return nil
}
