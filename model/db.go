package model

import (
  "github.com/jinzhu/gorm"
  _ "github.com/mattn/go-sqlite3"
)

func Migration() {
  db := GetDB()
  db.AutoMigrate(&User{})
  db.AutoMigrate(&Park{})
}

func GetDB() *gorm.DB {
  db, err := gorm.Open("sqlite3", "test.sqlite3")
  if err != nil {
    panic("failed to connect database\n")
  }
  return db
}
