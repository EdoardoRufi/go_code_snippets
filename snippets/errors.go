package snippets

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

type User struct{ ID int }

func repoGet(id int) (User, error) {
	if id == 0 {
		return User{}, ErrNotFound
	}
	return User{ID: id}, nil
}

func LoadUser(id int) (User, error) {
	u, err := repoGet(id)
	if err != nil {
		// %w doesn't change the error. -> errors.Is recognizes ErrNotFound. errors.Is walks the wrap chain to find ErrNotFound
		// %%v changes it. -> errors.Is doesn't recognizes it
		// return User{}, fmt.Errorf("load user id=%d: %w", id, err) // 404 is
		// return User{}, fmt.Errorf("load user id=%d: %v", id, err) // 500 is
		return User{}, fmt.Errorf("validate: %w", &ValidationError{Field: "email"}) // 400 as
	}
	return u, nil
}

type ValidationError struct {
	Field string
}

func (e *ValidationError) Error() string {
	return "invalid " + e.Field
}

func Error() {
	_, err := LoadUser(0)
	// errors.Is → “Is this error (or any wrapped one) equal to a specific error?”
	// errors.As → “Is there an error of this specific TYPE inside the chain?”
	if errors.Is(err, ErrNotFound) {
		fmt.Println("HTTP 404")
		return
	}
	var ve *ValidationError
	if errors.As(err, &ve) {
		fmt.Println("HTTP 400")
		return
	}
	fmt.Println("HTTP 500:", err)
}
