package utils

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   error  `json:"error"`
}

func (s Status) IsSuccess() bool {
	return s.Error == nil
}

func (s Status) IsError() bool {
	return s.Error != nil
}

func (s Status) NewStatus(code int, message string, err error) Status {
	return Status{Code: code, Message: message, Error: err}
}
