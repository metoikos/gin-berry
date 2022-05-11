package models

import (
	"gin-berry/db"
	"gin-berry/models/paging"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name     string    `gorm:"type:varchar(255)"`
	Email    string    `gorm:"type:varchar(255);index:;unique"`
	Password string    `gorm:"type:varchar(255)"`
}

// Create creates a user from the given pointer
func (u *User) Create() error {
	return db.DB.Create(u).Error
}

// Get returns the user with the given ID
func (u *User) Get() error {
	return db.DB.First(u).Error
}

// Update updates the user with the given pointer
func (u *User) Update() error {
	//return db.DB.Model(&u).Omit("ID", "Password", "CreatedAt", "UpdatedAt", "DeletedAt").Updates(&u).Error
	return db.DB.Model(&u).Select("Name", "Email").Updates(u).Error
}

// SetPassword sets the user's password
func (u *User) SetPassword(password []byte) error {
	p, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	return db.DB.Model(&u).Update("Password", p).Error
}

func (u *User) CheckPassword(password []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), password)
	if err != nil {
		return false
	}
	return true
}

// Delete deletes the user with the given pointer
func (u *User) Delete() error {
	return db.DB.Delete(&u).Error
}

// GetUsers returns all users
func (u User) GetUsers(page int64, limit int64) ([]User, paging.Paging) {
	var users []User
	var count int64

	// count the total number of items in the DB
	db.DB.Model(&User{}).Count(&count)

	// calculate the paging with the given numbers
	pagination := paging.New(page, limit, count)

	// filter the record with the right offset and limit
	db.DB.Find(&users).Order("created_at desc").Offset(int(pagination.Offset)).Limit(int(limit))

	return users, pagination
}
