package message

type Response struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"msg,omitempty"`
	Data    *Data  `json:"data,omitempty"`
}
