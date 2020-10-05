package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"html"
	"strings"
	"time"
)

type User struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:256" json:"name"`
	Family    string    `gorm:"size:256" json:"family"`
	Mobile    string    `gorm:"size:256" json:"mobile"`
	Email     string    `gorm:"size:256" json:"email"`
	Username  string    `gorm:"size:256" json:"username"`
	Password  string    `gorm:"size:256" json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (u *User) prepare() {
	u.ID = 0
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Family = html.EscapeString(strings.TrimSpace(u.Family))
	u.Mobile = html.EscapeString(strings.TrimSpace(u.Mobile))
	u.Email = html.EscapeString(strings.TrimSpace(strings.ToLower(u.Email)))
	u.Username = html.EscapeString(strings.TrimSpace(strings.ToLower(u.Username)))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) GetAllUsers(db *gorm.DB) (*[]User, error) {
	var err error

	users := []User{}

	err = db.Debug().Model(&User{}).Find(&users).Error

	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}
