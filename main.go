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
	r.MaxMultipartMemory = 32 << 20 // 32MB

	r.POST("/upload", server.UploadHandler)
	r.GET("/download/:type/:filename", server.DownloadHandler)

	log.Println("서버가 8080 포트에서 시작되었습니다...")
	r.Run(":8080")
}
