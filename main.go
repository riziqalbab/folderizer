package main

import (
	"context"
	"folderizer/action"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Commands: []*cli.Command{
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "get file list in the working directory or specified directory",
				Action:  action.ListDir,
			},
			{
				Name:    "set",
				Aliases: []string{"st"},
				Usage:   "",
			},
			{
				Name:    "start",
				Aliases: []string{"s"},
				Usage:   "Creates folders based on file types and moves files into the appropriate folders.",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "folder_prefix",
						Aliases: []string{"f", "fnp"},
						Usage:   "Add a first name to each folder, for example: --folder_prefix new_folder",
					},
					&cli.StringFlag{
						Name:    "folder_suffix",
						Aliases: []string{"s", "snp"},
						Usage:   "Add a name to the folder to be created, for example: --folder_suffix new_folder",
					},
				},
				Action: action.Start,
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

}
