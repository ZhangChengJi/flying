package config

type JWT struct {
	SigningKey  string `mapstructure:"signing-key" json:"signingKey" yaml:"signing-key"`
	ExpiresTime int64  `mapstructure:"expires-time" json:"expiresTime" yaml:"expires-time"`
	BufferTime  int64  `mapstructure:"buffer-time" json:"bufferTime" yaml:"buffer-time"`
	PrivateKey  string `mapstructure:"privateKey" json:"privateKey" yaml:"privateKey"`
	PublicKey   string `mapstructure:"publicKey" json:"publicKey" yaml:"publicKey"`
}
