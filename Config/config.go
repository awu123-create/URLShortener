package Config

import "os"

type Config struct {
	AppEnv  string
	AppPort string
	BaseUrl string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBCharset  string
}

func LoadConfig() *Config {
	return &Config{
		AppEnv:  os.Getenv("APP_ENV"),
		AppPort: os.Getenv("APP_PORT"),
		BaseUrl: os.Getenv("BASE_URL"),

		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBCharset:  os.Getenv("DB_CHARSET"),
	}
}
