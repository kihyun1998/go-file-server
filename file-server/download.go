package fileserver

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func (s *FileServer) DownloadHandler(c *gin.Context) {
	fileType := c.Param("type")
	filename := c.Param("filename")

	if !isValidFileType(fileType) {
		c.JSON(http.StatusBadRequest, Response{
			IsOk:  false,
			Error: "Invalid file type",
		})
		return
	}

	filePath := filepath.Join(s.uploadPath, fileType, filename)

	if !isValidPath(s.uploadPath, filePath) {
		c.JSON(http.StatusBadRequest, Response{
			IsOk:  false,
			Error: "Invalid file path",
		})
		return
	}

	// 파일 존재 여부 확인
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, Response{
			IsOk:  false,
			Error: "File not found",
		})
		return
	}

	c.File(filePath)
}
