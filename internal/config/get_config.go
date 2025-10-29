package config

import (
	"fmt"
	"log"

	"github.com/wb-go/wbf/config"
)

func GetConfig() (*Config, error) {
	cfg := config.New()
	cfg.SetDefault("HTTP_PORT", "8080")
	cfg.SetDefault("METRICS", false)
	cfg.SetDefault("PPROF", false)

	if err := cfg.LoadEnvFiles(".env"); err != nil {
		log.Printf("config.Load: %v. Продолжаем с дефолтными значениями", err)
	}
	cfg.EnableEnv("")

	var c Config
	if err := cfg.Unmarshal(&c); err != nil {
		return nil, fmt.Errorf("config.Unmarshal: %v", err)
	}

	return &c, nil
}
