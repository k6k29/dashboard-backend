package dockerCloud

type TestConnRequest struct {
	Host      string `json:"host"`
	Port      int    `json:"port"`
	UseTLS    bool   `json:"use_tls"`
	TLSCaCert string `json:"tls_ca_cert"`
	TLSCert   string `json:"tls_cert"`
	TLSKey    string `json:"tls_key"`
}
