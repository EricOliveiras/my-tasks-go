package builders

import "github/ericoliveiras/basic-crud-go/internal/models"

type UserBuilder struct {
	ID        string
	firstname string
	lastname  string
	email     string
	password  string
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{}
}

func (userBuilder *UserBuilder) SetID(id string) *UserBuilder {
	userBuilder.ID = id
	return userBuilder
}

func (userBuilder *UserBuilder) SetFirstname(firstname string) *UserBuilder {
	userBuilder.firstname = firstname
	return userBuilder
}

func (userBuilder *UserBuilder) SetLastname(lastname string) *UserBuilder {
	userBuilder.lastname = lastname
	return userBuilder
}

func (userBuilder *UserBuilder) SetEmail(email string) *UserBuilder {
	userBuilder.email = email
	return userBuilder
}

func (userBuilder *UserBuilder) SetPassword(password string) *UserBuilder {
	userBuilder.password = password
	return userBuilder
}

func (userBuilder *UserBuilder) Build() models.User {
	user := models.User{
		ID:        userBuilder.ID,
		FirstName: userBuilder.firstname,
		LastName:  userBuilder.lastname,
		Email:     userBuilder.email,
		Password:  userBuilder.password,
	}

	return user
}
