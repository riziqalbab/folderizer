package utils

import (
	"folderizer/typestructs"
	"strings"
)

func GetCategory(fileFormat string) string {

	categories := typestructs.Categories

	fileFormat = strings.TrimPrefix(fileFormat, ".")

	for _, category := range categories {
		for _, format := range category.Format {
			if format == fileFormat {
				return category.Name
			}
		}
	}

	return "other"
}
