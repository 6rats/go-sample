package controller

import (
  "github.com/gin-gonic/gin"
  "local.packages/model"
)

type UserController struct{}

func (uc UserController) Index(c *gin.Context) {
  db := model.GetDB()
  var parks []model.Park
  db.Find(&parks)
  c.JSON(200, parks)
}
