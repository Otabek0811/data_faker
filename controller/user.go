package controller

import (
	"app/models"
	"strings"

	"github.com/bxcodec/faker/v4"
)

func GenerateUser(count int) []models.User {
	var users []models.User
	for count >= 0 {
		users = append(users, models.User{
			Name:        faker.Name(),
			PhoneNumber: faker.Phonenumber(),
		})
		count--
	}

	return users
}

func SearchPhoneNumber(u []models.User, n string) []models.User {
	var usrs []models.User
	for _, val := range u {
		if strings.Contains(val.PhoneNumber, n) {
			usrs=append(usrs, val)
			return usrs
		}
	}
	return nil 
}
