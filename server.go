package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)

func clockHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, `
        <!DOCTYPE html>
        <html>
        <body>
            It's %d o'clock now.
        </body>
        </html>
    `, time.Now().Hour())
}

func main() {
    // ハンドラーを登録
    http.HandleFunc("/clock", clockHandler)
    // 先ほどの静的ファイルハンドラー
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("/etc/gohttpserver/doc"))))
    // サーバーを起動
    log.Fatal(http.ListenAndServe(":8080", nil))
}