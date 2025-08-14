package exceptions

import "fmt"

type InvalidCredentialsError struct {
	Reason string
}

func (i *InvalidCredentialsError) Error() string {
	return fmt.Sprintf(i.Reason)
}
