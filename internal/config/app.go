package config

type App struct {
	// 应用标识
	Id string `default:"${ID=.}" json:"id,omitempty"`
	// 密钥
	Secret string `default:"${SECRET}" json:"secret,omitempty"`
}
