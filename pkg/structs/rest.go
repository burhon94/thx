package structs

type ResponseAPI struct {
	Code    int         `json:"code"`
	Payload interface{} `json:"payload"`
	Info    string      `json:"info"`
	Errors  error       `json:"errors"`
}
