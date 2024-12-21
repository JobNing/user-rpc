package model

import (
	"github.com/JobNing/corehub/mysql"
	"gorm.io/gorm"
)

type Type struct {
	Model
	Name string `json:"name" gorm:"column:name;type:varchar(255)"`
	Icon string `json:"icon" gorm:"column:icon;type:char(11)"`
	Lv   int64  `json:"lv" gorm:"column:lv;type:tinyint(1)"`
	Pid  int64  `json:"pid" gorm:"column:pid;type:int(11)"`
}

func (m *Type) Create() (info *Type, err error) {
	return m, databases.WithClient(func(db *gorm.DB) error {
		return db.Create(m).Error
	})
}

func (m *Type) Update(id int64) (info *Type, err error) {
	return m, databases.WithClient(func(db *gorm.DB) error {
		return db.Model(m).Updates(m).Where("id = ?", id).Error
	})
}

func (m *Type) GetTypeByID(id int64) (info *Type, err error) {
	err = databases.WithClient(func(db *gorm.DB) error {
		return db.Where("id = ?", id).Find(&info).Error
	})
	return
}

func (m *Type) GetTypeListByPID(id int64) (infos []*Type, err error) {
	err = databases.WithClient(func(db *gorm.DB) error {
		return db.Where("pid = ?", id).Find(&infos).Error
	})
	return
}

func (m *Type) GetTypeListByPIDs(ids []int64) (infos []*Type, err error) {
	err = databases.WithClient(func(db *gorm.DB) error {
		return db.Where("pid in(?)", ids).Find(&infos).Error
	})
	return
}
