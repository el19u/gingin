package initializers

import(
  "fmt"
  "log"

  "gingin/models"

  "gorm.io/gorm"
  "gorm.io/driver/postgres"
  "github.com/spf13/viper"
)


var db *gorm.DB

func InitDB() {
  host := viper.GetString("database.host")
  user := viper.GetString("database.user")
  dbname := viper.GetString("database.name")
  port := viper.GetInt("database.port")
  sslmode := viper.GetString("database.sslmode")
  timezone := viper.GetString("database.timezone")

  dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%d sslmode=%s TimeZone=%s", host, user, dbname, port, sslmode, timezone)

  var err error
  db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    log.Fatal(err)
  }

  db.AutoMigrate(&models.Post{})
}
