package fileserver

import (
	"path/filepath"
)

type FileServer struct {
	uploadPath string
}

func NewFileServer(uploadPath string) (*FileServer, error) {
	absPath, err := filepath.Abs(uploadPath)
	if err != nil {
		return nil, err
	}

	if err := setupDirectory(absPath); err != nil {
		return nil, err
	}

	return &FileServer{uploadPath: absPath}, nil
}
