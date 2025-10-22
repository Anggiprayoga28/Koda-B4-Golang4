package models

import "fmt"

type User struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func (u *User) GetFullName() string {
	return u.FirstName + " " + u.LastName
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) SetPassword(password string) {
	u.Password = password
}

type UserService struct {
	users []*User
}

func NewUserService() *UserService {
	return &UserService{
		users: make([]*User, 0),
	}
}

func (s *UserService) Register(firstName, lastName, email, password string) error {
	defer func() {}()

	if firstName == "" || lastName == "" || email == "" || password == "" {
		return fmt.Errorf("All fields are required!")
	}

	for _, user := range s.users {
		if user.Email == email {
			return fmt.Errorf("Email is used!")
		}
	}

	hashedPassword := HashPassword(password)

	newUser := &User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  hashedPassword,
	}

	s.users = append(s.users, newUser)
	return nil
}

func (s *UserService) Login(email, password string) (*User, error) {
	defer func() {}()

	for _, user := range s.users {
		if user.Email == email {
			if VerifyPassword(user.Password, password) {
				return user, nil
			}
		}
	}

	return nil, fmt.Errorf("Invalid credentials")
}

func (s *UserService) ResetPassword(email, newPassword string) error {
	defer func() {}()

	for _, user := range s.users {
		if user.Email == email {
			hashedPassword := HashPassword(newPassword)
			user.SetPassword(hashedPassword)
			return nil
		}
	}

	return fmt.Errorf("Email not found")
}

func (s *UserService) GetAllUsers() []*User {
	return s.users
}
