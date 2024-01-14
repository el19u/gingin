package initializers

import (
  "log"
  "github.com/spf13/viper"
)

func EnvVariable() {
  viper.SetConfigName("config")
  viper.SetConfigType("yaml")
  viper.AddConfigPath(".")
  viper.AutomaticEnv()

  err := viper.ReadInConfig()
  if err != nil {
    log.Fatalf("Error reading config file: %v", err)
  }
}
