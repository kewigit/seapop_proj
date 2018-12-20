package main

import (
    "github.com/codegangsta/martini-contrib/render"
    "github.com/go-martini/martini"
    "html/template"
    "log"
    "net/http"
)

func htmlHandler0(w http.ResponseWriter, r *http.Request) {
    // テンプレートをパース
    t := template.Must(template.ParseFiles("/Go/bin/Go/src/sipo_porj/view/message.tpl"))

    str := "SeaPop"

    // テンプレートを描画
    if err := t.ExecuteTemplate(w, "message.tpl", str); err != nil {
        log.Fatal(err)
    }
}

func formPage(ren render.Render) {
    ren.HTML(200, "form", nil) 
}

func postPage(ren render.Render, req *http.Request) {
    req.ParseForm() // このメソッドを実行するとreq.Formに値が格納されます
    username := req.Form["Username"] // POSTされたデータを配列(スライス)で取得
    fmt.Println(username) // コンソールに値を表示
    ren.HTML(200, "hoge", nil)
  }

func main() {
    http.HandleFunc("/msg", htmlHandler0)
    m := martini.Classic()
    m.Use(render.Renderer())  // htmlテンプレートのレンダラーの指定
    m.Get("/", formPage) // // このページでフォームを表示します
    m.Post("/post_page", postPage) // POST後のページ
    m.Run()
    // サーバーを起動
    http.ListenAndServe(":8080", nil)
}