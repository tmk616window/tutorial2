package config

import "github.com/kelseyhightower/envconfig"

type Cfg struct {
    Env      string `required:"true" envconfig:"ENV"`
	PORT     string `required:"true" envconfig:"PORT"`
	Database
	GCP
}

func NewConfig() (*Cfg, error) {
	c := &Cfg{}
	if err := envconfig.Process("", c); err != nil {
		return nil, err
	}
	
	return c, nil
}

func (c *Cfg) IsLocal() bool {
	return c.Env == "local"
}
