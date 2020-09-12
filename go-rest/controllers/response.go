package controllers

// response http response type
type response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

// newResponse is response struct
func newResponse(status int, message string, result interface{}) *response {
	return &response{status, message, result}
}
