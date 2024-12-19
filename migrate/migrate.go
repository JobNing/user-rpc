package migrate

import (
	"github.com/JobNing/corehub/mysql"
	model2 "github.com/JobNing/user-rpc/internal/model"
	"gorm.io/gorm"
)

func AutoMigrate() error {
	return databases.WithClient(func(db *gorm.DB) error {
		return db.AutoMigrate(
			&model2.User{},
			&model2.Role{},
		)
	})
}
