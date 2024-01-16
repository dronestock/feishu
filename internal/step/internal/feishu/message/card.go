package message

type Card struct {
	Id       string         `json:"template_id,omitempty"`
	Variable map[string]any `json:"template_variable,omitempty"`
}
