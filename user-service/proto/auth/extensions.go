package auth

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

// BeforeCreate - hook to create hashed id
func (model *User) BeforeCreate(scope *gorm.Scope) error {

	newUUID := uuid.Must(uuid.NewV4())

	return scope.SetColumn("Id", newUUID.String())
}
