package utils

import (
	"github.com/go-playground/validator/v10"
	"log"
	"testing"
)

var validate = validator.New()

type User struct {
	Name     string `validate:"required" json:"name" msg_required:"Name is required!"`
	Email    string `validate:"required,email" json:"email" msg_required:"Email is required!"`
	Password string `validate:"required,max=64,min=6" json:"password" msg_required:"Please enter your password!" msg_min:"Please enter minimum {{.Param}} characters!" msg_max:"You can enter maximum {{.Param}} characters!"`
}

func TestErrorMessages(t *testing.T) {
	user := User{
		Name:     "John Doe",
		Email:    "test@test.com",
		Password: "asasd",
	}
	err := validate.Struct(user)

	apiErrors, _ := BuildAPiError(user, err)
	if len(apiErrors) != 1 {
		t.Fatalf("Invalid error output. It should be 1 error but got %v", len(apiErrors))
	}

	msg := apiErrors[0].Msg
	if msg != "Please enter minimum 6 characters!" {
		t.Fatalf("Invalid error output. The message should be 'Please enter minimum 6 characters!' but got %v", msg)
	}
}

func TestMultipleMessages(t *testing.T) {
	user := User{Email: "a"}
	err := validate.Struct(user)
	log.Println(err)

	apiErrors, _ := BuildAPiError(user, err)
	if len(apiErrors) != 3 {
		t.Fatalf("Invalid error output. It should be 1 error but got %v", len(apiErrors))
	}

	msg := apiErrors[0].Msg
	if msg != "Name is required!" {
		t.Fatalf("Invalid error output. The message should be 'Name is required!' but got %v", msg)
	}
}
