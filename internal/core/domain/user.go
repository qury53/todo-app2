package domain

import (
	"fmt"
	"net/mail"
	"regexp"
	"unicode/utf8"

	core_errors "github.com/qury53/todo-app2/internal/core/errors"
)

type User struct {
	ID          int
	Version     int
	Username    string
	FullName    string
	Password    string
	PhoneNumber *string
	Email       string
}

type UserPatch struct {
	FullName    Nullable[string]
	Password    Nullable[string]
	PhoneNumber Nullable[string]
	Email       Nullable[string]
}

func NewUser(
	id int,
	version int,
	username string,
	fullName string,
	password string,
	phoneNumber *string,
	email string,
) User {
	return User{
		ID:          id,
		Version:     version,
		Username:    username,
		FullName:    fullName,
		Password:    password,
		PhoneNumber: phoneNumber,
		Email:       email,
	}
}

func NewUserPatch(
	fullName Nullable[string],
	password Nullable[string],
	phoneNumber Nullable[string],
	email Nullable[string],
) UserPatch {
	return UserPatch{
		FullName:    fullName,
		Password:    password,
		PhoneNumber: phoneNumber,
		Email:       email,
	}
}

func NewUserUninitialized(
	username string,
	fullName string,
	password string,
	phoneNumber *string,
	email string,
) User {
	return User{
		ID:          userIDUnitialized,
		Version:     userVersionUninitialized,
		Username:    username,
		FullName:    fullName,
		Password:    password,
		PhoneNumber: phoneNumber,
		Email:       email,
	}
}

func (u *User) Validate() error {
	usernameLength := utf8.RuneCountInString(u.FullName)
	if usernameLength < 3 || usernameLength > 100 {
		return fmt.Errorf("validation error, username must be between 3 and 100: %w", core_errors.ErrInvalidArgument)
	}

	fullNameLength := utf8.RuneCountInString(u.FullName)
	if fullNameLength < 3 || fullNameLength > 100 {
		return fmt.Errorf("validation error, fullname must be between 3 and 100: %w", core_errors.ErrInvalidArgument)
	}

	passwordLength := utf8.RuneCountInString(u.Password)
	if passwordLength < 6 || passwordLength > 100 {
		return fmt.Errorf("validation error, password must be between 6 and 100: %w", core_errors.ErrInvalidArgument)
	}

	phoneNumber := utf8.RuneCountInString(*u.PhoneNumber)
	if phoneNumber < 10 || phoneNumber > 15 {
		return fmt.Errorf("validation error, phone number must be between 10 and 15: %w", core_errors.ErrInvalidArgument)
	}

	re := regexp.MustCompile(`^\+[0-9]+$`)

	if !re.MatchString(*u.PhoneNumber) {
		return fmt.Errorf("invalid phone number format: %s, err: %w", *u.PhoneNumber, core_errors.ErrInvalidArgument)
	}

	emailLength := utf8.RuneCountInString(u.Email)
	if emailLength > 255 {
		return fmt.Errorf("validation error, email length must be less or equal 255: %w", core_errors.ErrInvalidArgument)
	}

	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return fmt.Errorf("invalid email format: %w", core_errors.ErrInvalidArgument)
	}

	return nil
}

func (u *User) ApplyPatch(userPatch UserPatch) error {
	tmp := *u

	if userPatch.FullName.Set {
		tmp.FullName = *userPatch.FullName.Value
	}

	if userPatch.Password.Set {
		tmp.Password = *userPatch.Password.Value
	}

	if userPatch.PhoneNumber.Set {
		tmp.PhoneNumber = userPatch.PhoneNumber.Value
	}

	if userPatch.Email.Set {
		tmp.Email = *userPatch.Email.Value
	}

	*u = tmp

	return nil
}

func (u *UserPatch) Validate() error {
	if u.FullName.Set && u.FullName.Value == nil {
		return fmt.Errorf("fullname cannot be set to null: %w", core_errors.ErrInvalidArgument)
	}

	if u.Password.Set && u.Password.Value == nil {
		return fmt.Errorf("password cannot be set to null: %w", core_errors.ErrInvalidArgument)
	}

	if u.Email.Set && u.Email.Value == nil {
		return fmt.Errorf("email cannot be set to null: %w", core_errors.ErrInvalidArgument)
	}

	return nil
}
