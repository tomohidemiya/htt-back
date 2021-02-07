package user

import (
    "github.com/jinzhu/gorm"
)

type User struct {
    gorm.Model
    Email string
    LastName string
    FirstName string
    Tel string
}
