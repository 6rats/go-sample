module github.com/6rats/go-sample

go 1.13

require github.com/gin-gonic/gin v1.5.0

require github.com/jinzhu/gorm v1.9.11

require local.packages/controller v0.0.0

replace local.packages/model => ./model

require local.packages/model v0.0.0

replace local.packages/controller => ./controller
