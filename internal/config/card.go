package config

type Card struct {
	// 成功模板
	Success string `default:"${CARD_SUCCESS=ctp_AAiUZPR1RCxe}" json:"success,omitempty"`
	// 失败模板
	Failure string `default:"${CARD_FAILURE=ctp_AAiUZPR1RCxe}" json:"failure,omitempty"`
}
