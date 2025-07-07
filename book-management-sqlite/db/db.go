package db

import (
	"book-management-sqlite/config"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	// 连接 SQLite 数据库
	var err error
	DB, err = sql.Open("sqlite3", config.Cfg.DBFilePath)
	if err != nil {
		log.Fatalf("无法打开数据库文件: %v", err)
	}

	// 创建表
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		author TEXT NOT NULL,
		isbn TEXT UNIQUE NOT NULL,
		price REAL NOT NULL
	);`
	_, err = DB.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("创建表失败: %v", err)
	}

	log.Println("数据库初始化完成")
}
