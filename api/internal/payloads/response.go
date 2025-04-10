package payloads

type Response struct {
	Code       int    `json:"code"`
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	Data       any    `json:"data,omitempty"`
	Exceptions any    `json:"exceptions,omitempty"`
}
