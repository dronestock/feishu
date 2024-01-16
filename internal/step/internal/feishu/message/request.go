package message

type Request struct {
	Receive string `json:"receive_id,omitempty"`
	Type    string `json:"msg_type,omitempty"`
	Content string `json:"content,omitempty"`
	Id      string `json:"uuid,omitempty"`
}
