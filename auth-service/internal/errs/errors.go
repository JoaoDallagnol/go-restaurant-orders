package errs

import "fmt"

type UserNotFoundError struct {
	UserId uint
}

func (e *UserNotFoundError) Error() string {
	return fmt.Sprintf("user with Id %d not found", e.UserId)
}
