package app

// Confirmation generic request confirmation.
type Confirmation struct {
	Message string `json:"message"` // confirmation message
}

// ErrorResponse standard error response model.
type ErrorResponse struct {
	ID     string                   `json:"id,omitempty" xml:"id" form:"id"`                           // ID is the unique error instance identifier.
	Code   string                   `json:"code,omitempty" xml:"code" form:"code"`                     // Code identifies the class of errors.
	Status int                      `json:"status" xml:"status" form:"status"`                         // Status is the HTTP status code used by responses that cary the error.
	Detail string                   `json:"detail" xml:"detail" form:"detail"`                         // Detail describes the specific error occurrence.
	Meta   []map[string]interface{} `json:"meta,omitempty" xml:"meta,omitempty" form:"meta,omitempty"` // Meta contains additional key/value pairs useful to clients.
}
