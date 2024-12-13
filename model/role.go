package model

import (
	databases "github.com/JobNing/corehub/mysql"
	"gorm.io/gorm"
)

type Role struct {
	Model
	Name string `json:"name" gorm:"column:name;type:varchar(255)"`
}

func (m *Role) Create() (info *Role, err error) {
	return m, databases.WithClient(func(db *gorm.DB) error {
		return db.Create(m).Error
	})
}

func (m *Role) Update(id int64) (info *Role, err error) {
	return m, databases.WithClient(func(db *gorm.DB) error {
		return db.Model(m).Updates(m).Where("id = ?", id).Error
	})
}

func (m *Role) Get(id int64) (info *Role, err error) {
	return m, databases.WithClient(func(db *gorm.DB) error {
		return db.First(m, id).Error
	})
}

func (m *Role) SearchByName(name string) (infos []*Role, err error) {
	return infos, databases.WithClient(func(db *gorm.DB) error {
		return db.Where("name like ?", "%"+name+"%").Find(&infos).Error
	})
}
