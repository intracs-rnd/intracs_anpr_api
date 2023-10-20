package env

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error while reading env file. %s", err)
	}
}

func Get(key string) string {
	value, ok := viper.Get(key).(string)
	if !ok {
		fmt.Println("Invalid key assertion. '%s'", key)
	}

	return value
}
