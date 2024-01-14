package main

import(
  "fmt"
  "log"
  "net/http"

  "gingin/initializers"

  "github.com/gin-gonic/gin"
  "github.com/spf13/viper"
)

func init() {
  initializers.EnvVariable()
  initializers.InitDB()
}

func main() {
  router := gin.Default()

  router.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"data": "Hello World"})
  })

  port := viper.GetInt("server.port")
  addr := fmt.Sprintf(":%d", port)
  if err := router.Run(addr); err != nil {
    log.Fatal(err)
  }
}
