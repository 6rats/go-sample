package model

import (
  "github.com/jinzhu/gorm"
)

type User struct {
  gorm.Model
  Name string
  Age int
}

type Park struct {
  gorm.Model
  Name string
  Lat float32
  Lon float32
}

