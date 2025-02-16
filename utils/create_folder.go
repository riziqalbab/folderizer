package utils

import (
	"fmt"
	"os"
	"path"
)

func CreateFolder(paths string, name string) {

	if _, err := os.Stat(path.Join(paths, name)); os.IsNotExist(err) {
		os.Mkdir(path.Join(paths, name), 0777)
		fmt.Print(err)
	}

}
