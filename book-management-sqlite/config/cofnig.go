package config

import (
	"os"
)

type Config struct {
	Port       string
	DBFilePath string // SQLite 数据库文件路径
}

var Cfg Config

func LoadConfig() {
	Cfg.Port = getEnv("PORT", "8080")
	Cfg.DBFilePath = getEnv("DB_FILE_PATH", "./data/books.db")
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return fallback
}
