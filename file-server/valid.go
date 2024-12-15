package fileserver

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
	"unicode"
)

const (
	MaxFileSize       = 100 * 1024 * 1024 // 100MB
	MaxFileNameLength = 255
)

var (
	allowedImageTypes = map[string][]string{
		".jpg":  {"image/jpeg"},
		".jpeg": {"image/jpeg"},
		".png":  {"image/png"},
		".gif":  {"image/gif"},
	}

	allowedVideoTypes = map[string][]string{
		".mp4": {"video/mp4"},
		".avi": {"video/x-msvideo"},
		".mov": {"video/quicktime"},
		".wmv": {"video/x-ms-wmv"},
	}
)

// 안전한 파일명 생성
func generateSafeFileName(originalName string) string {
	// 랜덤 접두사 생성
	randomBytes := make([]byte, 16)
	rand.Read(randomBytes)
	prefix := hex.EncodeToString(randomBytes)

	// 파일 확장자 추출 및 검증
	ext := strings.ToLower(filepath.Ext(originalName))
	if ext == "" {
		ext = ".bin"
	}

	// 파일명 길이 제한
	if len(originalName) > MaxFileNameLength {
		originalName = originalName[:MaxFileNameLength-len(ext)]
	}

	// 특수문자 제거 및 공백을 언더스코어로 변경
	safeName := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || r == '-' || r == '_' {
			return r
		}
		if unicode.IsSpace(r) {
			return '_'
		}
		return -1
	}, originalName)

	return prefix + "_" + safeName + ext
}

// MIME 타입 검증
func validateFileContent(file multipart.File, filename string) (string, error) {
	// 파일 버퍼 읽기
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		return "", err
	}

	// 파일 포인터를 처음으로 되돌림
	file.Seek(0, 0)

	// MIME 타입 감지
	contentType := http.DetectContentType(buffer)
	ext := strings.ToLower(filepath.Ext(filename))

	// 이미지 검증
	if allowedMimes, ok := allowedImageTypes[ext]; ok {
		for _, mime := range allowedMimes {
			if contentType == mime {
				return "images", nil
			}
		}
	}

	// 비디오 검증
	if allowedMimes, ok := allowedVideoTypes[ext]; ok {
		for _, mime := range allowedMimes {
			if contentType == mime {
				return "videos", nil
			}
		}
	}

	return "", fmt.Errorf("unsupported file type: %s", contentType)
}

// 경로 검증 강화
func isValidPath(basePath, targetPath string) bool {
	// 경로 정규화
	cleanBase := filepath.Clean(basePath)
	cleanTarget := filepath.Clean(targetPath)

	// 절대 경로로 변환
	absBase, err := filepath.Abs(cleanBase)
	if err != nil {
		return false
	}
	absTarget, err := filepath.Abs(cleanTarget)
	if err != nil {
		return false
	}

	// 경로 비교
	rel, err := filepath.Rel(absBase, absTarget)
	if err != nil {
		return false
	}

	// 상위 디렉토리 참조 검사
	return !strings.Contains(rel, "..")
}
