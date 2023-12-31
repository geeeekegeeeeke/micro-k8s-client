package response

type FileInfo struct {
	//fileutil.FileInfo
}

type UploadInfo struct {
	Name      string `json:"name"`
	Size      int    `json:"size"`
	CreatedAt string `json:"createdAt"`
}

type FileTree struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	Path     string     `json:"path"`
	Children []FileTree `json:"children"`
}

type DirSizeRes struct {
	Size float64 `json:"size" validate:"required"`
}

type FileProcessKeys struct {
	Keys []string `json:"keys"`
}

type FileWgetRes struct {
	Key string `json:"key"`
}
