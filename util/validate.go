// Provides validation functionality.

package util

import "github.com/go-playground/validator"

// validate is the global validator.
var validate = validator.New()

// Validate validates a given structure to see if
// it matches its definition.
//
// It returns a nice output.
func Validate(v any) string {
	err := validate.Struct(v)

	var msg string

	errs := err.(validator.ValidationErrors)
	for i, err := range errs {
		msg += err.Tag() + " " + err.Field()

		// Add a comma separator on items that are
		// not the last.
		if i < len(errs)-1 {
			msg += ", "
		}
	}

	return msg
}
