package message

type Content struct {
	Type string `json:"type,omitempty"`
	Data *Card  `json:"data,omitempty"`
}
