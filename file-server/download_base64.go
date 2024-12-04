package fileserver

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func (s *FileServer) DownloadBase64Handler(c *gin.Context) {
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

	// 파일 읽기
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			c.JSON(http.StatusNotFound, Response{
				IsOk:  false,
				Error: "File not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, Response{
				IsOk:  false,
				Error: "Failed to read file",
			})
		}
		return
	}

	// Content-Type 설정
	contentType := "application/octet-stream"
	if isImage(filename) {
		if filepath.Ext(filename) == ".png" {
			contentType = "image/png"
		} else if filepath.Ext(filename) == ".jpg" || filepath.Ext(filename) == ".jpeg" {
			contentType = "image/jpeg"
		} else if filepath.Ext(filename) == ".gif" {
			contentType = "image/gif"
		}
	} else if isVideo(filename) {
		if filepath.Ext(filename) == ".mp4" {
			contentType = "video/mp4"
		} else if filepath.Ext(filename) == ".avi" {
			contentType = "video/x-msvideo"
		}
	}

	// 헤더 설정
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", contentType)

	// 파일 데이터를 직접 전송
	if _, err := c.Writer.Write(fileData); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			IsOk:  false,
			Error: "Failed to send file",
		})
		return
	}
}
