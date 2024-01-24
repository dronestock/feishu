package config

type Card struct {
	// 模板
	Template string `default:"${CARD_TEMPLATE}" json:"template,omitempty"`
}
