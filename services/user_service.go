package services

import (
	"kanbanmusume_ura/db"
	"kanbanmusume_ura/models"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}
type User models.User

func (_ UserService) FindByName(name string) (User, error) {
	db := db.Connect()
	var user User
	if err := db.Where("name = ?", name).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (_ UserService) ConvertHash(password string) string {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}

	return string(hashedPass)
}

func (_ UserService) IsSamePassword(savedPasswordHash string, requestedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(savedPasswordHash), []byte(requestedPassword))

	return err
}
