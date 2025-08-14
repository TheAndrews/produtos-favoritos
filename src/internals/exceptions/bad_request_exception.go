package exceptions

import "fmt"

type BadRequestError struct {
	Reason string
}

func (i *BadRequestError) Error() string {
	return fmt.Sprintf(i.Reason)
}
