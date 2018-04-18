package auth

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

// BeforeCreate - hook to create hashed id
func (model *User) BeforeCreate(scope *gorm.Scope) error {

	newUUID, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("Id", newUUID.String())
}
