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
	// path := path.Base("/")

	// file_list, _ := os.ReadDir(path)

	cmd := &cli.Command{
		Commands: []*cli.Command{
			{
				Name:    "getFileList",
				Aliases: []string{"a"},
				Usage:   "add a task to the list",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					path_args := cmd.Args().Get(0)
					file_list, err := utils.GetFileList(path_args)

					if err != nil {
						fmt.Println(err)
					}

					for _, file := range file_list {
						fmt.Println(file.Name())
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
