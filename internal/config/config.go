package config

type Config struct {
	HTTPPort string `mapstructure:"HTTP_PORT"`
	Metrics  bool   `mapstructure:"METRICS"`
	Pprof    bool   `mapstructure:"PPROF"`
}
