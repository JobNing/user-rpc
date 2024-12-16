package migrate

import (
	"github.com/JobNing/corehub/mysql"
	"github.com/JobNing/user-rpc/model"
	"gorm.io/gorm"
)

func AutoMigrate() error {
	return databases.WithClient(func(db *gorm.DB) error {
		return db.AutoMigrate(
			&model.User{},
			&model.Role{},
		)
	})
}
