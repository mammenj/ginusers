package models

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"time"
)

// User "Object
type User struct {
	ID        int       `json:"id"`
	UID       uuid.UUID `json:"uuid"`
	Username  string    `json:"username" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//BeforeCreate BeforeCreate
func (u *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())
	scope.SetColumn("UID", uuid.NewV4().String())
	return nil
}

//BeforeUpdate BeforeUpdate
func (u *User) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}
