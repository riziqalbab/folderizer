package main

import (
	"context"
	"fmt"
	"folderizer/utils"
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
				Action: func(ctx context.Context, cmd *cli.Command) error {
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
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

}
