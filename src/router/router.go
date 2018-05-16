package router

import (
  "net/http"
  martini "github.com/go-martini/martini"
  render "github.com/martini-contrib/render"
)

func user(router martini.Router) {

  router.Any("", func(res http.ResponseWriter, req *http.Request) string {
    return "user"
  })

  router.Get("/:id", func(params martini.Params, r render.Render){
    // return `{"id":"` + params["id"] + `"}`
    // data := make(map[string]string)
    // data["id"] = params["id"]
    r.JSON(200, params)
  })

}



func Router(m *martini.ClassicMartini) {
  // 首页
  m.Get("/", func(r render.Render){
    r.HTML(200, "index", nil)
  })
  // 用户组
  m.Group("/user", func(r martini.Router) {
    user(r)
  })
}
