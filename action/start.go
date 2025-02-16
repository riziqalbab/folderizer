package action

import (
	"context"
	"fmt"
	"folderizer/utils"
	"os"

	"github.com/urfave/cli/v3"
)

func Start(ctx context.Context, cmd *cli.Command) error {

	pathsArg := cmd.Args().Get(0)
	workingDirectoryPath, err := os.Getwd()

	if err != nil {
		panic(fmt.Errorf("error getting working directory: %w", err))
	}

	if pathsArg != "" {
		fmt.Printf("Organizing %s\n", pathsArg)
	} else {
		fmt.Printf("Organizing %s\n", workingDirectoryPath)
		utils.Organize(workingDirectoryPath)
	}

	return nil
}
