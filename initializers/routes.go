package initializers

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

// SetupRoutes 設置所有的路由
func SetupRoutes(router *gin.Engine) {
  router.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"data": "Hello World"})
  })
}
