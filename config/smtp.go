package config

type SmtpConfig struct {
	From     string `mapstructure:"from"`
	Host     string `mapstructure:"host"`
	Port     int16  `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}
