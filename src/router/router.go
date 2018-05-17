package router

import (
  "router/book"
  martini "github.com/go-martini/martini"
  render "github.com/martini-contrib/render"
)


func Router(m *martini.ClassicMartini) {
  // 首页
  m.Get("/", func(r render.Render){
    // r.HTML(200, "index", nil)
    data := make(map[string]string)
    // data["title"] = "goweb"
    r.HTML(200, "index", data)
  })
  // 用户组
  m.Group("/book", func(r martini.Router) {
    book.Book(r)
  })
}
