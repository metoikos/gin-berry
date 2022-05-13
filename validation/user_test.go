package validation

import (
	//"github.com/fatih/structtag"
	"github.com/go-playground/validator/v10"
	"testing"
)

var validate = validator.New()

func TestUser(t *testing.T) {
	user := User{
		Name:     "John Doe",
		Email:    "test@test.com",
		Password: "asdasdsad",
	}
	err := validate.Struct(user)
	if err != nil {
		t.Fatalf("Invalid validation. %v should pass all the rules.", user)
	}
}

func TestUserFail(t *testing.T) {
	user := User{
		Name:     "John Doe",
		Email:    "test@test.com",
		Password: "asasd",
	}
	err := validate.Struct(user)

	if err == nil {
		t.Fatalf("Invalid validation. %v should fail on password validation", user)
	}
}

func TestUserFailMinLength(t *testing.T) {
	user := User{
		Name:     "John Doe",
		Email:    "test@test.com",
		Password: "asdas",
	}
	err := validate.Struct(user)
	if err == nil {
		t.Fatalf("Invalid validation. %v should fail on password validation", user)
	}
}

func TestUserFailMaxLength(t *testing.T) {
	user := User{
		Name:     "John Doe",
		Email:    "test@test.com",
		Password: "asdasadadsad3plasdlas;dlas;das;dlas;dlasd;lasd;la;l1';3l4';l23rm,as",
	}
	err := validate.Struct(user)
	if err == nil {
		t.Fatalf("Invalid validation. %v should fail on password validation", user)
	}
}
