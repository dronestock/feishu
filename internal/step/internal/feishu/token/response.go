package token

type Response struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"msg,omitempty"`
	Token   string `json:"tenant_access_token,omitempty"`
	Expire  int    `json:"expire,omitempty"`
}
