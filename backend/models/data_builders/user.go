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

func (b *UserBuilder) WithUserID(id uint64) *UserBuilder {
	b.User.UserID = id
	return b
}

func (b *UserBuilder) WithName(name string) *UserBuilder {
	b.User.Name = name
	return b
}

func (b *UserBuilder) WithEmail(email string) *UserBuilder {
	b.User.Email = email
	return b
}

func (b *UserBuilder) WithPassword(password string) *UserBuilder {
	b.User.Password = password
	return b
}

func (b *UserBuilder) WithType(t models.UserType) *UserBuilder {
	b.User.Type = t
	return b
}

func (b *UserBuilder) Build() *models.User {
	return b.User
}
