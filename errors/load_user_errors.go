package errors

import (
	"errors"
	"fmt"
	"go_code_snippets/errors/secure"
)

type User struct{ ID int }

// service
func LoadUser(id int) (*User, error) {
	if id == 0 {
		return nil, &secure.SafeError{
			Code:     "FETCH_ERROR",
			UserMsg:  "Unable to retrieve user profile.",
			Internal: NotFoundErr, // Stored for logs, hidden from Unwrap logic if needed
		}
	}
	if id < 0 {
		return nil, NewValidationError(
			"Invalid input",
			ValidationErr,
			map[string]any{
				"username": id,
			},
		)
	}
	return &User{ID: id}, nil
}

// controller
func ControllerLoadUser(id int) (*User, error) {
	u, err := LoadUser(id)
	if err != nil {

		var safeErr *secure.SafeError
		// errors.Is → “Is this error (or any wrapped one) equal to a specific error?”
		// errors.As → “Is there an error of this specific TYPE inside the chain?”
		if errors.As(err, &safeErr) {
			fmt.Println("Internal log: ", safeErr.LogString())
			fmt.Println("external log returned: ", safeErr.UserMsg)
			return &User{}, fmt.Errorf("external log returned: %s", safeErr.UserMsg)
		} else {
			return &User{}, fmt.Errorf("generic error calling load user id=%d: %w", id, err)
		}
	}
	return u, nil
}
