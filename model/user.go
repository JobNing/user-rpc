package model

import (
	"errors"
	"github.com/JobNing/corehub/mysql"
	"gorm.io/gorm"
)

type User struct {
	Model
	Username string `json:"username" gorm:"column:username;type:varchar(255)"`
	Phone    string `json:"phone" gorm:"column:phone;type:char(11);unique"`
	Age      int64  `json:"age" gorm:"column:age;type:int(11)"`
	Sex      int64  `json:"sex" gorm:"column:sex;type:tinyint(1)"`
}

func (m *User) Create() (info *User, err error) {
	return m, databases.WithClient(func(db *gorm.DB) error {
		return db.Create(m).Error
	})
}

func (m *User) Update(id int64) (info *User, err error) {
	return m, databases.WithClient(func(db *gorm.DB) error {
		return db.Model(m).Updates(m).Where("id = ?", id).Error
	})
}

func (m *User) Get(id int64) (info *User, err error) {
	return m, databases.WithClient(func(db *gorm.DB) error {
		return db.First(m, id).Error
	})
}

func (m *User) GetByPhone(phone string) (info *User, err error) {
	return m, databases.WithClient(func(db *gorm.DB) error {
		err = db.Where("phone=?", phone).First(m).Error
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return nil
	})
}
