package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	// 测试数据库连接
	if err := db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("Database connected")
	return db, nil
}
