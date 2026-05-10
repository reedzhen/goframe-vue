package payment

import "goframe-vben/internal/model/dto"

var config *dto.PayConfig

func SetConfig(c *dto.PayConfig) {
	config = c
}

func GetConfig() *dto.PayConfig {
	return config
}
