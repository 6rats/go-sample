package main

import (
  "local.packages/model"
  "github.com/gin-gonic/gin"
  "strconv"
)

func main() {
  router := gin.Default()

  model.Migration()

  router.GET("/", func(c *gin.Context) {
    c.String(200, "Hello,World!")
  })

  router.GET("/parks", func(c *gin.Context) {
    pref := c.DefaultQuery("pref", "Tokyo")//パラメータが未指定の時に効く
    if pref == "" {
      pref = "Empty"
    }
    size := c.Query("size")
    i, _ := strconv.Atoi(size)//文字列→Int
    c.String(200, "Park Pref = %s Size = %d", pref, i)
  })

  router.GET("/parks/:id", func(c *gin.Context) {
    c.String(200, "Park ID = " + c.Param("id"))
  })

  router.Run()
}
