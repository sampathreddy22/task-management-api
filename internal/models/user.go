package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// @Description User input model for creation and updates
type UserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"required,oneof=user admin"`
}

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key"`
	Email        string    `gorm:"not null;unique"`
	PasswordHash string    `gorm:"not null" json:"-"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Role         string
	Tasks        []Task `gorm:"foreignKey:UserID"`
}

// function to hash the password
func (u *User) hashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(hashedPassword)
	return nil
}
func (u *User) FromUserInput(userInput UserInput) {
	u.Email = userInput.Email
	u.Role = userInput.Role
	u.ID = uuid.New()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	u.hashPassword()
}
