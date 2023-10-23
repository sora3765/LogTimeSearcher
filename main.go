package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// 日本時間のタイムゾーンを設定
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		http.Error(w, "タイムゾーンの設定エラー", http.StatusInternalServerError)
		return
	}

	currentTime := time.Now().In(jst).Format("02/Jan/2006:15:04:05")
	html := fmt.Sprintf("<html><body>現在の時刻 (日本時間): %s</body></html>", currentTime)
	fmt.Fprint(w, html)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Webサーバーを開始します...")
	http.ListenAndServe(":8080", nil)
}
