package config;

import (
	"os"
    "reflect"
	"github.com/joho/godotenv"
	"fmt"
)

type Config struct {
	HTTP_PORT    string
	VERSION     string
	SERVICE_NAME string
	JWT_SECRET string
}


var cfg Config;

func LoadConfig() {
	
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	v := reflect.ValueOf(&cfg).Elem()

	envList := []string{"HTTP_PORT", "VERSION", "SERVICE_NAME", "JWT_SECRET"}
	for _, key := range envList {
		if val := os.Getenv(key); val != "" {
			field := v.FieldByName(key)
			if field.IsValid() && field.CanSet() {
				field.SetString(val)
			}
		}
	}

}


func GetConfig() Config {

	
	return cfg;
}

