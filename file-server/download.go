package fileserver

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func (s *FileServer) DownloadHandler(c *gin.Context) {
	fileType := c.Param("type")
	filename := c.Param("filename")

	if !isValidFileType(fileType) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "잘못된 파일 타입",
		})
		return
	}

	filePath := filepath.Join(s.uploadPath, fileType, filename)

	if !isValidPath(s.uploadPath, filePath) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "잘못된 파일 경로",
		})
		return
	}

	c.File(filePath)
}
