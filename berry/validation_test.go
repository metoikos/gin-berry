package berry

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

type QueryParams struct {
	Page  uint8 `form:"page" json:"page"  validate:"number,min=1" msg_min:"Invalid page!" msg_number:"Page must be a number!"`
	Limit uint8 `form:"limit" json:"limit" validate:"number,oneof=10 25 50" msg_number:"Limit must be a number!" msg_oneof:"Invalid limit value!"`
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

	apiErrors, _ := BuildAPiError(user, err)
	if len(apiErrors) != 3 {
		t.Fatalf("Invalid error output. It should be 1 error but got %v", len(apiErrors))
	}

	msg := apiErrors[0].Msg
	if msg != "Name is required!" {
		t.Fatalf("Invalid error output. The message should be 'Name is required!' but got %v", msg)
	}
}

func TestQueryParams(t *testing.T) {
	q := QueryParams{
		Page:  0,
		Limit: 12,
	}
	err := validate.Struct(q)
	apiErrors, _ := BuildAPiError(q, err)
	if len(apiErrors) != 2 {
		t.Fatalf("Invalid error output. It should be 2 error but got %v", len(apiErrors))
	}
	//
	msg := apiErrors[0].Msg
	if msg != "Invalid page!" {
		t.Fatalf("Invalid error output. The message should be 'Invalid page!' but got %v", msg)
	}
}

func TestQueryParams2(t *testing.T) {
	q := QueryParams{
		Page:  1,
		Limit: 10,
	}
	err := validate.Struct(q)

	apiErrors, _ := BuildAPiError(q, err)
	if len(apiErrors) != 0 {
		t.Fatalf("It should pass the validation but got %v", len(apiErrors))
	}
}

type Params struct {
	Querystring interface{}
}

type Config struct {
	params Params
}

func getConfig() Config {
	return Config{
		params: Params{
			Querystring: QueryParams{
				Page:  1,
				Limit: 12,
			},
		},
	}
}

func parse(p Params) {
	err := validate.Struct(p.Querystring)
	log.Println("Err", err)
}

func TestQueryParamsEmpty(t *testing.T) {
	conf := getConfig()
	parse(conf.params)
	//q := QueryParams{}
	//err := validate.Struct(q)
	//
	//log.Println("ERR XXX", err)
}
