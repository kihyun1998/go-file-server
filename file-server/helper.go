package fileserver

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type ResponseData struct {
	Path     string `json:"path,omitempty"`
	Filename string `json:"filename,omitempty"`
	FileType string `json:"fileType,omitempty"`
}

type Response struct {
	IsOk    bool         `json:"isOk"`
	Message string       `json:"message,omitempty"`
	Error   string       `json:"error,omitempty"`
	Data    ResponseData `json:"data,omitempty"`
}

func setupDirectory(path string) error {
	directories := []string{
		filepath.Join(path, "images"),
		filepath.Join(path, "videos"),
	}
	// Windows와 Linux 모두를 위한 권한 설정
	var perm os.FileMode = 0755
	if runtime.GOOS == "windows" {
		perm = 0666
	}
	for _, dir := range directories {
		if err := os.MkdirAll(dir, perm); err != nil {
			return err
		}
	}
	return nil
}

func getUploadPath(basePath, fileType string) string {
	// 운영체제에 맞는 경로 구분자 사용
	path := filepath.Join(basePath, fileType)

	// Windows와 Linux 모두를 위한 권한 설정
	var perm os.FileMode = 0755
	if runtime.GOOS == "windows" {
		perm = 0666
	}

	if err := os.MkdirAll(path, perm); err != nil {
		return basePath
	}

	return path
}

// 파일 권한 설정 함수
func setFilePermissions(filename string) error {
	if runtime.GOOS == "windows" {
		return os.Chmod(filename, 0666)
	}
	return os.Chmod(filename, 0644)
}

// 파일 타입 검사 헬퍼 함수들
func isImage(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif"
}

func isVideo(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".mp4" || ext == ".avi" || ext == ".mov" || ext == ".wmv"
}

func isValidFileType(fileType string) bool {
	return fileType == "images" || fileType == "videos" || fileType == "others"
}
