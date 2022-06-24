package srenvelope

// Status defines the structure to manipulate a
// response status information structure.
type Status struct {
	Success bool            `json:"success" xml:"success"`
	Errors  StatusErrorList `json:"errors" xml:"errors"`
}

// NewStatus instantiates a new request result status structure.
func NewStatus() *Status {
	return &Status{
		Success: true,
		Errors:  StatusErrorList{},
	}
}

// AddError append a new error to the status error list
func (s *Status) AddError(err *StatusError) *Status {
	s.Success = false
	s.Errors = append(s.Errors, err)
	return s
}

// SetService assign a service code to all stored errors.
func (s *Status) SetService(val int) *Status {
	for i := range s.Errors {
		s.Errors[i] = s.Errors[i].SetService(val)
	}
	return s
}

// SetEndpoint assign an endpoint code to all stored errors.
func (s *Status) SetEndpoint(val int) *Status {
	for i := range s.Errors {
		s.Errors[i] = s.Errors[i].SetEndpoint(val)
	}
	return s
}
