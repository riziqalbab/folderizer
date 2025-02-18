package action

import (
	"context"
	"errors"
	"fmt"

	"github.com/urfave/cli/v3"
)

func SetAction(ctx context.Context, cmd *cli.Command) error {
	// pathsArg := cmd.Args().Get(0)
	// workingDirectoryPath, err := os.Getwd()

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
	organizeSet(ctx, cmd)
	return nil
}

func organizeSet(ctx context.Context, cmd *cli.Command) {

}
