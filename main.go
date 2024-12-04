package main

import (
	"log"

	"github.com/gin-gonic/gin"
	fileserver "github.com/kihyun1998/go-file-server/file-server"
)

func main() {
	server, err := fileserver.NewFileServer("./uploads")
	if err != nil {
		log.Fatalf("파일 서버 초기화 실패: %v", err)
	}

	r := gin.Default()
	r.MaxMultipartMemory = 0 // 파일 크기 제한 해제

	r.POST("/upload", server.UploadHandler)
	r.POST("/upload/base64", server.UploadBase64Handler)
	r.GET("/download/:type/:filename", server.DownloadHandler)
	r.GET("/download/base64/:type/:filename", server.DownloadBase64Handler)

	log.Println("서버가 8080 포트에서 시작되었습니다...")
	r.Run(":8080")
}
