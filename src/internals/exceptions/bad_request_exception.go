package exceptions

type BadRequestError struct {
	Reason string
}

func (i *BadRequestError) Error() string {
	return i.Reason
}
