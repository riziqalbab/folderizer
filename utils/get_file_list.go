package utils

import (
	"os"
)

func GetFileList(path string) ([]os.DirEntry, error) {

	file_list, error := os.ReadDir(path)

	return file_list, error

}
