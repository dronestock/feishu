package config

type Card struct {
	// 成功模板
	Success string `default:"${CARD_SUCCESS=ctp_AAikgQRejIVp}" json:"success,omitempty"`
	// 失败模板
	Failure string `default:"${CARD_FAILURE=ctp_AAikmSMerakX}" json:"failure,omitempty"`
}
