package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// DetailedMessage はメッセージの詳細情報を保持します
type DetailedMessage struct {
	Action      string `json:"action"`
	Description string `json:"description"`
	UserID      string `json:"user_id,omitempty"`
}

// LogEntry は Cloud Logging の構造化ログのフォーマットを定義します
type LogEntry struct {
	Severity  string          `json:"severity"`
	Message   DetailedMessage `json:"message"`
	Timestamp time.Time       `json:"timestamp"`
}

// structuredLog は構造化ログを標準出力に書き込みます
func structuredLog(severity string, message DetailedMessage) {
	entry := LogEntry{
		Severity:  severity,
		Message:   message,
		Timestamp: time.Now(),
	}

	jsonEntry, err := json.Marshal(entry)
	if err != nil {
		log.Printf("ログのJSONエンコードに失敗しました: %v", err)
		return
	}

	fmt.Println(string(jsonEntry))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		structuredLog("INFO", DetailedMessage{
			Action:      "request_received",
			Description: "新しいリクエストを受信しました",
			UserID:      r.Header.Get("X-User-ID"),
		})

		fmt.Fprintf(w, "Hello, World!")

		structuredLog("INFO", DetailedMessage{
			Action:      "response_sent",
			Description: "レスポンスを送信しました",
		})
	})

	structuredLog("INFO", DetailedMessage{
		Action:      "server_start",
		Description: "サーバーを起動します",
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
