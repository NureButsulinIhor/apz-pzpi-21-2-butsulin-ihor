package models

import "gorm.io/gorm"

type UserType string

const (
	WorkerType  UserType = "worker"
	ManagerType UserType = "manager"
	AdminType   UserType = "admin"
)

type User struct {
	gorm.Model
	Email   string `json:"email" gorm:"unique"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
	Type    UserType
}

func NewUserFromClaims(claims map[string]interface{}) (user User, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	id := uint(claims["id"].(float64))
	email := claims["email"].(string)
	name := claims["name"].(string)
	picture := claims["picture"].(string)
	userType := UserType(claims["type"].(string))

	return User{
		Model:   gorm.Model{ID: id},
		Email:   email,
		Name:    name,
		Picture: picture,
		Type:    userType,
	}, err
}

func (u User) GetClaims() map[string]interface{} {
	return map[string]interface{}{
		"email":   u.Email,
		"name":    u.Name,
		"picture": u.Picture,
		"type":    u.Type,
	}
}
