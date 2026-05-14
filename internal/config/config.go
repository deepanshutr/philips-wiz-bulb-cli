package config

import "os"

type Config struct {
	CoreURL string
}

func Load() Config {
	c := Config{CoreURL: os.Getenv("PHILIPS_WIZ_BULB_CORE_URL")}
	if c.CoreURL == "" {
		c.CoreURL = "http://127.0.0.1:8766"
	}
	return c
}
