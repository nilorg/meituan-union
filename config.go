package union

// Config ...
type Config struct {
	BaseURL      string `json:"base_url"`
	AppKey       string `json:"app_key"`
	SignatureKey string `json:"signature_key"`
}

// NewConfig ...
func NewConfig(appKey, signatureKey string) *Config {
	return &Config{
		BaseURL:      BaseURL,
		AppKey:       appKey,
		SignatureKey: signatureKey,
	}
}
