package main

import(
  "fmt"
  "log"

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

  initializers.SetupRoutes(router)

  port := viper.GetInt("server.port")
  addr := fmt.Sprintf(":%d", port)
  if err := router.Run(addr); err != nil {
    log.Fatal(err)
  }
}
