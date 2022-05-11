package validation

type User struct {
	Name     string `validate:"required" json:"name"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,max=64,min=6" json:"password"`
}
