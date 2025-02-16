package utils

import (
	"fmt"
	"os"
)

func MoveFile(src string, dst string) {
	err := os.Rename(src, dst)
	if err != nil {
		fmt.Println(err)
	}
}
