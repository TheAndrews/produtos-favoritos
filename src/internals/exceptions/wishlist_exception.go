package exceptions

import "fmt"

type AlreadyWishlistedErr struct {
	Reason string
}

func (i *AlreadyWishlistedErr) Error() string {
	return fmt.Sprintf("%s", i.Reason)
}
