package config

type Card struct {
	// 成功模板
	// https://open.feishu.cn/tool/cardbuilder?from=op_develop_app&templateId=ctp_AAikgQRejIVp
	Success string `default:"${CARD_SUCCESS=ctp_AAikgQRejIVp}" json:"success,omitempty"`
	// 失败模板
	// https://open.feishu.cn/tool/cardbuilder?from=op_develop_app&templateId=ctp_AAikmSMerakX
	Failure string `default:"${CARD_FAILURE=ctp_AAikmSMerakX}" json:"failure,omitempty"`
}
