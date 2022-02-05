package config

type JwtConfig struct {
	Secret              string `mapstructure:"secret"`
	AccessTokenExpDays  int    `mapstructure:"access_token_exp_days"`
	RefreshTokenExpDays int    `mapstructure:"refresh_token_exp_days"`
}
