package main

import (
  "local.packages/model"
  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()

  model.Migration()

  router.GET("/", func(c *gin.Context) {
    c.String(200, "Hello,World!")
  })

  router.Run()
}
