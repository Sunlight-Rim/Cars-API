package response

// General response struct.
type Response struct {
	Response any           `json:"response"`
	Error    ErrorResponse `json:"error"`
}

// General error response struct.
type ErrorResponse struct {
	Code    uint32 `json:"code"`
	Message string `json:"message"`
}

type Signup struct {
	ID uint64
}
