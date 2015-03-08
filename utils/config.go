package utils

import (
	"code.google.com/p/gcfg"
)

type Config struct {
	Section Section
}

type Section struct {
	Name string
}

func New() *Config {
	cfg := &Config{
		Section: Section{
			Name: "hel",
		},
	}
	return cfg
}

func (c *Config) LoadConfig(path string) error {
	err := gcfg.ReadFileInto(c, path)
	return err
}
