package exceptions

import "fmt"

type NotFoundEntityError struct {
	Reason string
}

func (i *NotFoundEntityError) Error() string {
	return fmt.Sprintf("%s", i.Reason)
}
