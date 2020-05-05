package main

import (
  "local.packages/model"
  "local.packages/router"
  "github.com/gin-gonic/gin"
)

func main() {

  r := gin.Default()
  model.Migration()
  router.Routing(r)
  r.Run()

  /*
  router.GET("/parks", func(c *gin.Context) {
    pref := c.DefaultQuery("pref", "Tokyo")//パラメータが未指定の時に効く
    if pref == "" {
      pref = "Empty"
    }
    size := c.Query("size")
    i, _ := strconv.Atoi(size)//文字列→Int
    db := model.GetDB()
    var count int = 0
    db.Model(&model.Park{}).Count(&count)
    c.String(200, "Park Pref = %s Size = %d Count = %d", pref, i, count)
  })

  router.GET("/parks/:id", func(c *gin.Context) {
    c.String(200, "Park ID = " + c.Param("id"))
  })

  router.POST("/parks", func(c *gin.Context) {
    db := model.GetDB()
    db.Create(&model.Park{Name: "Test Park", Lat: 35.734621, Lon: 139.7098774})
    db.Close()
  })
  */
}
