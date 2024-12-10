package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	fileserver "github.com/kihyun1998/go-file-server/file-server"
)

func main() {
	server, err := fileserver.NewFileServer("./uploads")
	if err != nil {
		log.Fatalf("파일 서버 초기화 실패: %v", err)
	}

	r := gin.Default()

	// 정적 파일 제공을 위한 라우터 추가
	r.Static("/static", "./static")

	// HTML 템플릿 로드
	r.LoadHTMLGlob("templates/*")

	// CORS 설정
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept"}
	r.Use(cors.New(config))

	r.MaxMultipartMemory = 0 // 파일 크기 제한 해제

	r.POST("/upload", server.UploadHandler)
	r.POST("/upload/base64", server.UploadBase64Handler)
	r.GET("/download/:type/:filename", server.DownloadHandler)
	r.GET("/download/base64/:type/:filename", server.DownloadBase64Handler)

	r.GET("/page/download/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		c.HTML(http.StatusOK, "download.html", gin.H{
			"filename": filename,
		})
	})

	log.Println("서버가 8080 포트에서 시작되었습니다...")
	r.Run(":8080")
}
