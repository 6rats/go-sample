package controller

import (
  "github.com/gin-gonic/gin"
  "local.packages/model"
  "net/http"
  "log"
)

type ParkController struct{}

type IndexParams struct {
  Pref int `form:"pref" binding:"required,min=1,max=47"`
}

func (uc ParkController) Index(c *gin.Context) {

  var params IndexParams

  if err := c.ShouldBind(&params); err != nil {
    c.JSON(http.StatusBadRequest, err.Error())
  } else {

    log.Println(params.Pref)

    var parks []model.Park
    db := model.GetDB()
    db.Find(&parks)
    c.JSON(http.StatusOK, parks)
  }
}
