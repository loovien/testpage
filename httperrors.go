package testpage

type HttpError struct {
	Code    int
	Message string
}

func (e *HttpError) Error() string {
	return e.Message
}

type E404Error struct {
	*HttpError
}

func NewE404Error(e *E404Error) (e404error *E404Error) {
	return &E404Error{&HttpError{404, "page missing"}}
}
