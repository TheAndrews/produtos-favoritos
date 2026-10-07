package exceptions

type InvalidEntityError struct {
	Reason string
}

func (i *InvalidEntityError) Error() string {
	return i.Reason
}
