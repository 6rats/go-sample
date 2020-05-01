package model

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

func Migration() {
  db := GetDB()
  db.AutoMigrate(&User{})
  db.AutoMigrate(&Park{})
}

func GetDB() *gorm.DB {
  db, err := gorm.Open("mysql", "test_user:test_pass@(db)/parks?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    panic("failed to connect database\n")
  }
  return db
}
