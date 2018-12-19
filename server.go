package main

import (
    "html/template"
    "log"
    "net/http"
)

func htmlHandler0(w http.ResponseWriter, r *http.Request) {
    // テンプレートをパース
    t := template.Must(template.ParseFiles("bin/Go/src/sipo_porj/view/message.tpl"))

    str := "SeaPop Message"

    // テンプレートを描画
    if err := t.ExecuteTemplate(w, "message.tpl", str); err != nil {
        log.Fatal(err)
    }
}

func main() {
    http.HandleFunc("/msg", htmlHandler0)

    // サーバーを起動
    http.ListenAndServe(":8080", nil)
}