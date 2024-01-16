package token

type Request struct {
	Id     string `json:"app_id,omitempty"`
	Secret string `json:"app_secret,omitempty"`
}
