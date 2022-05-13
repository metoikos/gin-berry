package validation

type User struct {
	Name     string `validate:"required" json:"name" msg_required:"Name is required!"`
	Email    string `validate:"required,email" json:"email" msg_required:"Email is required!"`
	Password string `validate:"required,max=64,min=6" json:"password" msg_required:"Please enter your password!" msg_min:"Please enter minimum {{.Param}} characters!" msg_max:"You can enter maximum {{.Param}} characters!"`
}
