package fileserver

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

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

func isValidPath(basePath, targetPath string) bool {
	// 경로 정규화
	basePath = filepath.Clean(basePath)
	targetPath = filepath.Clean(targetPath)

	// 대소문자 구분 없이 검사 (Windows의 경우)
	if runtime.GOOS == "windows" {
		basePath = strings.ToLower(basePath)
		targetPath = strings.ToLower(targetPath)
	}

	return strings.HasPrefix(targetPath, basePath)
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
