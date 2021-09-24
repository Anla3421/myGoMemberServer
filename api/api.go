package api

import "github.com/gin-gonic/gin"

func Api() {
	r := gin.Default()
	Router(r)
	listenPort := "9000"
	r.Run(":" + listenPort)
}
