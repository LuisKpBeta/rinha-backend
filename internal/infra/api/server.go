package api

import "github.com/gin-gonic/gin"

func CreateHttpServer() *gin.Engine {
	r := gin.Default()
	return r
}
func StartHttpServer(server *gin.Engine) {
	server.Use(gin.Logger())
	server.Run(":9999")
}
