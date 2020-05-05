package router

import (
  "local.packages/controller"
  "github.com/gin-gonic/gin"
)

func Routing(e *gin.Engine) {

  e.GET("/", func(c *gin.Context) {
    c.String(200, "暫定")
  })

  parks := e.Group("/parks")
  {
    c := controller.ParkController{}
    parks.GET("", c.Index)
  }
}
