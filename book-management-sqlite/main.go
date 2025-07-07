package main

import (
	"book-management-sqlite/config"
	"book-management-sqlite/db"
	"book-management-sqlite/routes"
	"log"
)

func main() {
	config.LoadConfig()
	db.InitDB()

	r := routes.SetupRoutes()
	if err := r.Run(":" + config.Cfg.Port); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
