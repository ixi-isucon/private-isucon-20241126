package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	ID      int
	Mime    string
	Imgdata []byte
}

func main() {
	// データベース接続
	dsn := "isuconp:isuconp@tcp(192.168.1.11:3306)/isuconp"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("データベース接続エラー: %v", err)
	}
	defer db.Close()

	// 保存先ディレクトリを作成
	outputDir := "/home/isucon/private_isu/webapp/public/image"
	err = os.MkdirAll(outputDir, 0755)
	if err != nil {
		log.Fatalf("ディレクトリ作成エラー: %v", err)
	}

	// データベースから画像データを取得
	rows, err := db.Query("SELECT id, mime, imgdata FROM posts")
	if err != nil {
		log.Fatalf("クエリエラー: %v", err)
	}
	defer rows.Close()

	// 各画像をファイルに保存
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Mime, &post.Imgdata)
		if err != nil {
			log.Printf("行のスキャンエラー: %v", err)
			continue
		}

		// ファイル拡張子を決定
		ext := getFileExtension(post.Mime)
		if ext == "" {
			log.Printf("未対応のMIMEタイプ: %s (Post ID: %d)", post.Mime, post.ID)
			continue
		}

		// ファイルパスを作成
		filename := fmt.Sprintf("%d%s", post.ID, ext)
		filepath := filepath.Join(outputDir, filename)

		// ファイルに書き込み
		err = ioutil.WriteFile(filepath, post.Imgdata, 0644)
		if err != nil {
			log.Printf("ファイル書き込みエラー (Post ID: %d): %v", post.ID, err)
			continue
		}

		log.Printf("画像保存完了: %s", filepath)
	}

	log.Println("すべての画像を保存しました。")
}

// MIMEタイプに基づいてファイル拡張子を取得
func getFileExtension(mime string) string {
	switch mime {
	case "image/jpeg":
		return ".jpg"
	case "image/png":
		return ".png"
	case "image/gif":
		return ".gif"
	default:
		return ""
	}
}
