package models

import (
	"gin-berry/db"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name     string    `gorm:"type:varchar(255)"`
	Email    string    `gorm:"type:varchar(255)"`
	Password string    `gorm:"type:varchar(255)"`
}

func (u *User) Create() error {
	return db.DB.Create(u).Error
}

func (u *User) Get() error {
	return db.DB.First(u).Error
}

func (u *User) Update() error {
	//return db.DB.Model(&u).Omit("ID", "Password", "CreatedAt", "UpdatedAt", "DeletedAt").Updates(&u).Error
	return db.DB.Model(&u).Update("Name", "Email").Error
}

func (u *User) Delete() error {
	return db.DB.Delete(&u).Error
}

func (u User) GetUsers(page int64, limit int64) ([]User, *Paging) {
	var users []User
	var count int64

	// count the total number of items in the DB
	db.DB.Model(&User{}).Count(&count)

	// calculate the paging with the given numbers
	paging := BuildPaging(page, limit, count)

	// filter the record with the right offset and limit
	db.DB.Find(&users).Order("created_at desc").Offset(int(paging.offset)).Limit(int(limit))

	return users, paging
}
