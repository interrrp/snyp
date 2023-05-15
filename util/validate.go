// Provides validation functionality.

package util

import "github.com/go-playground/validator"

// validate is the global validator.
var validate = validator.New()

// Validate validates a given structure to see if
// it matches its definition.
func Validate(v any) error {
	return validate.Struct(v)
}
