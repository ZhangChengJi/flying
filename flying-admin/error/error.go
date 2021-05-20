package error

type FlyingError struct {
	s string
}

func (fe *FlyingError) Error() string {
	return fe.s
}
