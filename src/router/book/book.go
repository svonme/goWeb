package book

import (
  "fmt"
  // "time"
  martini "github.com/go-martini/martini"
  render "github.com/martini-contrib/render"
  // 数据库
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Users struct {
    gorm.Model
    Name string `gorm:"default:''"`
}

func Book(router martini.Router) {

  router.Any("", func(r render.Render) {
    db, err := gorm.Open("mysql", "")
    if err != nil {
      panic("failed to connect database")
    }
    // 函数结束时自动关闭数据库
    defer db.Close()

    db.AutoMigrate(&Users{})
    db.DB()
    //
    users := []Users{}
    //
    db.Find(&users)
    //
    // // for index, item := range users{
      fmt.Println(users)
    // // }
    //
    r.JSON(200, users)


    // user := Users{Name: "svon"}
    // db.Create(&user)
    // r.JSON(200, user)
  })

  router.Get("/:id", func(params martini.Params, r render.Render){
    r.JSON(200, params)
  })

}
