package payloads

type Response struct {
	Code       int         `json:"code"`
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Exceptions interface{} `json:"exceptions,omitempty"`
}
