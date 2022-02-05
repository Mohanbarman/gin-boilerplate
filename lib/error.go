package lib

type ServiceError struct {
	Code int
	P    interface{}
}

func Error(code int) *ServiceError {
	return &ServiceError{
		Code: code,
		P:    nil,
	}
}

func (s *ServiceError) Payload(p interface{}) *ServiceError {
	s.P = p
	return s
}
