package exceptions

type InvalidCredentialsError struct {
	Reason string
}

func (i *InvalidCredentialsError) Error() string {
	return i.Reason
}
