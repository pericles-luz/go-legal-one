package legal_one

import (
	"encoding/json"

	"github.com/pericles-luz/go-base/pkg/conf"
)

type Config struct {
	file     conf.ConfigBase
	User     string `json:"user"`
	Password string `json:"password"`
	LinkAuth string `json:"linkAuth"`
	LinkAPI  string `json:"linkAPI"`
}

func (c *Config) Load(file string) error {
	raw, err := c.file.ReadConfigurationFile(file)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(raw), c)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) GetConfig() map[string]interface{} {
	return map[string]interface{}{
		"DE_User":  c.User,
		"PW_Senha": c.Password,
		"LN_Auth":  c.LinkAuth,
		"LN_API":   c.LinkAPI,
	}
}

func NewConfig() *Config {
	return &Config{}
}
