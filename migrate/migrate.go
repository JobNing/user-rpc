package migrate

import (
	"github.com/JobNing/corehub/mysql"
	"gorm.io/gorm"
	"user-rpc/model"
)

func AutoMigrate() error {
	return databases.WithClient(func(db *gorm.DB) error {
		return db.AutoMigrate(
			&model.User{},
			&model.Role{},
		)
	})
}
