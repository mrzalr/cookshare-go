package models

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GORM HOOKS
// CREATE USER ID AND HASH PASSWORD
func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	fmt.Println("called")

	return u.HashPassword()
}

func (u *User) HashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hash)
	return nil
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func (u *User) Sanitize() {
	u.Password = ""
}

func (u *User) GenerateToken() (string, error) {
	jwtSigningMethod := jwt.SigningMethodHS256
	jwtMapClaims := jwt.MapClaims{"id": u.ID}

	token := jwt.NewWithClaims(jwtSigningMethod, jwtMapClaims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}

type UserWithToken struct {
	User
	Token string `json:"token"`
}
