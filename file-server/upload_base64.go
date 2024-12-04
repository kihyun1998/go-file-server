package fileserver

import (
	"encoding/base64"
	"io"
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

	fileType := "others"
	if isImage(req.Filename) {
		fileType = "images"
	} else if isVideo(req.Filename) {
		fileType = "videos"
	}

	uploadPath := getUploadPath(s.uploadPath, fileType)
	filename := filepath.Join(uploadPath, req.Filename)

	// 파일 생성
	file, err := os.Create(filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			IsOk:  false,
			Error: "Failed to create file",
		})
		return
	}
	defer file.Close()

	// 디코딩된 데이터를 파일에 쓰기
	if _, err := io.Copy(file, strings.NewReader(string(fileData))); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			IsOk:  false,
			Error: "Failed to save file",
		})
		return
	}

	if err := setFilePermissions(filename); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			IsOk:  false,
			Error: "Failed to set file permissions",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		IsOk:    true,
		Message: "File uploaded successfully",
		Data: ResponseData{
			Path:     filepath.Join(fileType, req.Filename),
			Filename: req.Filename,
			FileType: fileType,
		},
	})
}
