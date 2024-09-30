package data_builders

import "cookdroogers/models"

type UserBuilder struct {
	User *models.User
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{
		User: &models.User{
			UserID:   7,
			Name:     "uzi",
			Email:    "uzi@gmail.com",
			Password: "password",
			Type:     models.NonMemberUser,
		},
	}
}

func (b *UserBuilder) Build() *models.User {
	return b.User
}
