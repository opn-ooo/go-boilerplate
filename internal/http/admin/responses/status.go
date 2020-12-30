package responses

// Status ... common status response structure
type Status struct {
	Success bool   `json:"success"`
	Error   *Error `json:"error"`
}

// NewSuccessStatus ...
func NewSuccessStatus() *Status {
	status := &Status{
		Success: true,
	}
	return status
}

// NewErrorStatus ...
func NewErrorStatus(code string, message string) *Status {
	status := &Status{
		Success: false,
		Error: &Error{
			Code:    code,
			Message: message,
		},
	}

	return status
}

// NewServiceUnavailableErrorStatus ... Return when the service is unavailable
func NewServiceUnavailableErrorStatus() *Status {
	return NewErrorStatus(ServiceUnavailableError, "This service is temporary unavailable")
}

// NewNotImplementedErrorStatus ... Return when the API is not implemented yet ( only for development period )
func NewNotImplementedErrorStatus() *Status {
	return NewErrorStatus(NotImplementedError, "This API is not implemented yet")
}

// NewInvalidParameterErrorStatus ... Return parameter error when input parameter from the client doesn't satisfy the requirement
func NewInvalidParameterErrorStatus(err error) *Status {
	return NewErrorStatus(InvalidParameterError, "Invalid Parameter: "+err.Error())
}

// NewInternalServerError ... Return if internal server error happens
func NewInternalServerError(err error) *Status {
	return NewErrorStatus(InternalServerError, "Internal Server Error: "+err.Error())
}

// NewInvalidHeaderErrorStatus ... Return header error when request headers from the client doesn't satisfy the requirement
func NewInvalidHeaderErrorStatus(message string) *Status {
	return NewErrorStatus(InvalidParameterError, "Invalid Header: "+message)
}

// NewNotFoundError ... Return if specified object not found
func NewNotFoundError(name string) *Status {
	return NewErrorStatus(NotFoundError, "Not Found: "+name)
}
