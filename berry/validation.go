package berry

import (
	"bytes"
	"errors"
	"github.com/go-playground/validator/v10"
	"reflect"
	"text/template"
)

// ApiError inspired from this post https://stackoverflow.com/a/70072158/12078
type ApiError struct {
	Field string `json:"field"`
	Msg   string `json:"message"`
}

// BuildAPiError generates an ApiError slice from a validator.ValidationErrors.
// We return either a slice of ApiError or an error.
// it basically reflects the struct type and reads field tags. From the tags it looks for the
// `msg_${validation_rule}` tag to assume as the error message.
// If the tag is not found, then it will use the default error message.
func BuildAPiError(st interface{}, baseError error) ([]ApiError, error) {
	var ve validator.ValidationErrors
	if errors.As(baseError, &ve) {
		out := make([]ApiError, len(ve))
		// convert st pointer to a struct
		// get the struct type
		stType := reflect.TypeOf(st)
		el := stType.Elem()

		for i, err := range ve {
			// get the tags from the base struct
			field, _ := el.FieldByName(err.Field())
			//log.Println("field =>", err.Tag())

			// We expect the error messages to be in the "msg_TAG_NAME" tag.
			// Combine the failed validation field and add "_required" to reach to the validation message.
			tagName := "msg_" + err.Tag()
			// We can ignore the error here because it might not be there
			jsonTag := field.Tag.Get(tagName)
			// get the initial error message
			errorMessage := err.Error()
			//log.Println("jsonTag=>", jsonTag)
			//log.Println("=======")
			// if an error message is exist
			if len(jsonTag) > 0 {
				//errorMessage = jsonTag.Name
				buf := &bytes.Buffer{}
				// We allow the user to use a template to generate the error message
				templ := template.Must(template.New(tagName).Parse(jsonTag))
				// apply fields to the template
				if terr := templ.Execute(buf, map[string]interface{}{
					"Param": err.Param(), // modifier value on the tag max=6 => 6 is the param
					"Value": err.Value(), // Value of the field = "asdasd"
					"Type":  err.Type(),  // Type of the field = string
					"Field": err.Field(), // Field name = Password
				}); terr == nil {
					// if there's no error, then we can replace the error message with the template
					errorMessage = buf.String()
				}
			}
			// put error message to the output slice
			out[i] = ApiError{err.Field(), errorMessage}
		}

		return out, nil
	}

	return nil, baseError
}
