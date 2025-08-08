package exceptions

import "fmt"

type EmailAlreadyRegisteredErr struct {
	Reason string
}

func (i *EmailAlreadyRegisteredErr) Error() string {
	return fmt.Sprintf("%s", i.Reason)
}
