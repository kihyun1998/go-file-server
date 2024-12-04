package fileserver

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func (s *FileServer) UploadHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "fail to read file.",
		})
		return
	}

	// 파일 타입 확인
	fileType := "others"
	if isImage(file.Filename) {
		fileType = "images"
	} else if isVideo(file.Filename) {
		fileType = "videos"
	}

	// 저장 경로 생성
	uploadPath := getUploadPath(s.uploadPath, fileType)
	filename := filepath.Join(uploadPath, file.Filename)

	// 파일 저장
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "fail to save file.",
		})
		return
	}

	// 파일 권한 설정
	if err := setFilePermissions(filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "file to set file permissions.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"isOK":     true,
		"message":  "success file upload.",
		"filename": file.Filename,
		"path":     filepath.Join(fileType, file.Filename),
	})
}
