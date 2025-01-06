// By Saif Hamdan, Team Lead
// Date: 2025/1/6
package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

type Validate struct {
	*validator.Validate
}

func New() *Validate {
	// Create a new instance of validator
	v := validator.New()

	// Register custom validation rule
	v.RegisterValidation("username", usernameValidation)

	return &Validate{
		Validate: v,
	}
}

// Username validation rule
func usernameValidation(fl validator.FieldLevel) bool {
	username := fl.Field().String()
	// Define regex pattern for username
	pattern := "^[a-zA-Z0-9_]{3,20}$"
	// patte1rn := "^[A-Za-z][A-Za-z0-9_]{7,29}$"

	matched, _ := regexp.MatchString(pattern, username)
	return matched
}
