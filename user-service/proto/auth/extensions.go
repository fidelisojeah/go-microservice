package auth

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// BeforeCreate - hook to create hashed id
func (model *User) BeforeCreate(scope *gorm.Scope) error {

	newUUID, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("Id", newUUID.String())
}
