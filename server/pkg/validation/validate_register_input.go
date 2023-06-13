package validation

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

type RegisterInput struct {
	FirstName    string `validate:"required,alpha,min=3,max=50"`
	LastName     string `validate:"required,alpha,min=3,max=50"`
	Username     string `validate:"required,username,min=4,max=20"`
	EmailAddress string `validate:"required,email"`
	Password     string `validate:"required,password,min=8,max=100"`
}

// Custom validation error messages
var customFieldErrors = map[string]string{
	"RegisterInput.FirstName":    "First name is required, must be alphabetical, and between 3 to 50 characters.",
	"RegisterInput.LastName":     "Last name is required, must be alphabetical, and between 3 to 50 characters.",
	"RegisterInput.Username":     "Username is required, can only contain letters, numbers, underscores, and dashes, and must be between 4 to 20 characters.",
	"RegisterInput.EmailAddress": "A valid email address is required.",
	"RegisterInput.Password":     "Password is required, must be at least 8 characters long, and contain at least 1 lowercase letter, 1 uppercase letter, 1 number, and 1 symbol.",
}

func ValidateRegisterInput(registerInput *RegisterInput) error {

	validate := validator.New()

	// Add custom validation for password
	_ = validate.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()
		fmt.Println("Password: ", password)
		length := len(password) >= 8
		uppercase := regexp.MustCompile(`[A-Z]`).MatchString(password)
		lowercase := regexp.MustCompile(`[a-z]`).MatchString(password)
		number := regexp.MustCompile(`[0-9]`).MatchString(password)
		symbol := regexp.MustCompile(`[^\w\s]`).MatchString(password)
		fmt.Println(length, uppercase, lowercase, number, symbol)
		return length && uppercase && lowercase && number && symbol
	})

	// Add custom validation for username
	_ = validate.RegisterValidation("username", func(fl validator.FieldLevel) bool {
		return regexp.MustCompile(`^[a-zA-Z0-9_-]+$`).MatchString(fl.Field().String())
	})

	err := validate.Struct(registerInput)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		for _, err := range err.(validator.ValidationErrors) {
			return fmt.Errorf(customFieldErrors[err.Namespace()])
		}
	}

	// User input is valid
	return nil
}
