package typestructs

type FileCategory struct {
	Name   string
	Format []string
}

var Categories = []FileCategory{
	{
		Name:   "image",
		Format: []string{"jpg", "jpeg", "png", "gif", "bmp", "svg"},
	},
	{
		Name:   "video",
		Format: []string{"mp4", "mkv", "avi", "mov"},
	},
	{
		Name:   "audio",
		Format: []string{"mp3", "wav", "ogg", "flac"},
	},
	{
		Name:   "document",
		Format: []string{"doc", "docx", "pdf", "xls", "xlsx", "ppt", "pptx"},
	},
	{
		Name:   "archive",
		Format: []string{"zip", "tar", "gz", "7z", "rar"},
	},
	{
		Name:   "programs",
		Format: []string{"exe", "msi", "sh"},
	},
}
