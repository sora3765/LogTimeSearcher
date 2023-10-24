//go:build js && wasm
// +build js,wasm

package main

import (
	"fmt"
	"syscall/js"
	"time"
)

func main() {
	c := make(chan struct{}, 0)

	// ブラウザからのHTTPリクエストを処理するハンドラ
	httpHandler := js.FuncOf(func(this js.Value, p []js.Value) interface{} {
		// 日本時間のタイムゾーンを設定
		jst, err := time.LoadLocation("Asia/Tokyo")
		if err != nil {
			fmt.Println("タイムゾーンの設定エラー")
			return nil
		}

		// 現在の日本時間を取得
		currentTime := time.Now().In(jst).Format("02/Jan/2006:15:04:05")
		html := fmt.Sprintf("<html><body>現在の時刻 (日本時間): %s</body></html>", currentTime)

		// HTML要素に挿入
		document := js.Global().Get("document")
		outputDiv := document.Call("getElementById", "output")
		outputDiv.Set("innerHTML", html)

		return nil
	})

	// ルートURLにハンドラを設定
	js.Global().Set("httpHandler", httpHandler)

	fmt.Println("WebAssemblyアプリを開始します...")

	// WebAssemblyが終了しないようにチャネルを使ってブロック
	<-c
}
