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
				Name:    "getFileList",
				Aliases: []string{"a"},
				Usage:   "get file list in the working directory or specified directory",
				Action:  action.ListDir,
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

}
