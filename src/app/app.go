package app

import (
  "os"
  "router"
  martini "github.com/go-martini/martini"
  render "github.com/martini-contrib/render"
)

func config() {
  os.Setenv("PORT", "9000") //设置端口号
}

func Run() {
  // 配置初始化
  config()
  // 实列化 martini 对象
  m := martini.Classic()

  m.Use(render.Renderer(render.Options{
    Directory: "views",
    Layout: "layouts/layout",
    Extensions: []string{".tmpl", ".html"},
    Charset: "UTF-8",
    IndentJSON: true,
    IndentXML: true,
  }))

  // 构造路由
  router.Router(m)
  m.Run()
}
