package database

import (
	"fmt"
	"log"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// データベース接続を初期化する
func InitDB() error {
	// 環境変数から接続情報を取得
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// 接続文字列の構築
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName,
	)

	// データベースに接続
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("データベース接続エラー: %v", err)
		return err
	}

	DB = db
	log.Println("データベース接続成功")
	return nil
}

// データベース接続を閉じる
func CloseDB() {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			log.Printf("データベース接続のクローズ中にエラーが発生: %v", err)
			return
		}
		sqlDB.Close()
		log.Println("データベース接続を閉じました")
	}
}

