package fileserver

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func (s *FileServer) UploadHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			IsOk:  false,
			Error: "Failed to read file",
		})
		return
	}

	fileType := "others"
	if isImage(file.Filename) {
		fileType = "images"
	} else if isVideo(file.Filename) {
		fileType = "videos"
	}

	uploadPath := getUploadPath(s.uploadPath, fileType)
	filename := filepath.Join(uploadPath, file.Filename)

	if err := c.SaveUploadedFile(file, filename); err != nil {
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
			Path:     filepath.Join(fileType, file.Filename),
			Filename: file.Filename,
			FileType: fileType,
		},
	})
}
