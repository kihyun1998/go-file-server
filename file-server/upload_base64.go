package fileserver

import (
	"encoding/base64"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

type Base64UploadRequest struct {
	FileData string `json:"fileData"`
	Filename string `json:"filename"`
}

func (s *FileServer) UploadBase64Handler(c *gin.Context) {
	var req Base64UploadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			IsOk:  false,
			Error: "Failed to parse request",
		})
		return
	}

	if req.Filename == "" {
		c.JSON(http.StatusBadRequest, Response{
			IsOk:  false,
			Error: "Filename is required",
		})
		return
	}

	if req.FileData == "" {
		c.JSON(http.StatusBadRequest, Response{
			IsOk:  false,
			Error: "File data is required",
		})
		return
	}

	// base64 문자열에서 "data:image/jpeg;base64," 같은 프리픽스 제거
	cleanData := req.FileData
	if idx := strings.Index(cleanData, ","); idx != -1 {
		cleanData = cleanData[idx+1:]
	}

	// base64 디코딩
	fileData, err := base64.StdEncoding.DecodeString(cleanData)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			IsOk:  false,
			Error: "Invalid base64 data",
		})
		return
	}

	// 파일 크기 체크
	if len(fileData) > MaxFileSize {
		c.JSON(http.StatusBadRequest, Response{
			IsOk:  false,
			Error: "File size exceeds limit",
		})
		return
	}

	// 안전한 파일명 생성
	safeFilename := generateSafeFileName(req.Filename)

	// 임시 파일 생성하여 MIME 타입 체크
	tempFile, err := os.CreateTemp("", "upload-*")
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			IsOk:  false,
			Error: "Failed to process file",
		})
		return
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	if _, err := tempFile.Write(fileData); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			IsOk:  false,
			Error: "Failed to process file",
		})
		return
	}
	tempFile.Seek(0, 0)

	// 파일 타입 검증
	fileType, err := validateFileContent(tempFile, safeFilename)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			IsOk:  false,
			Error: err.Error(),
		})
		return
	}

	uploadPath := getUploadPath(s.uploadPath, fileType)
	filename := filepath.Join(uploadPath, safeFilename)

	// 경로 검증
	if !isValidPath(s.uploadPath, filename) {
		c.JSON(http.StatusBadRequest, Response{
			IsOk:  false,
			Error: "Invalid file path",
		})
		return
	}

	// 파일 저장
	if err := os.WriteFile(filename, fileData, 0644); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			IsOk:  false,
			Error: "Failed to save file",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		IsOk:    true,
		Message: "File uploaded successfully",
		Data: ResponseData{
			Path:     filepath.Join(fileType, safeFilename),
			Filename: safeFilename,
			FileType: fileType,
		},
	})
}
