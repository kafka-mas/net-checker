package readconf

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type YamlConf struct {
	Phone     string   `yaml:"phone"`
	Addresses []string `yaml:"addresses"`
}

func (c *YamlConf) ConfigRead(file string) error {
	data, err := os.ReadFile(file)
	if err != nil {
		err = fmt.Errorf("Error open file: %v", err)
		return err
	}
	err = yaml.Unmarshal(data, c)
	if err != nil {
		err = fmt.Errorf("Error parse YAML: %v", err)
		return err
	}

	return nil
}

func (c *YamlConf) ConfigGetPhone() string {
	return c.Phone
}

func (c *YamlConf) ConfigGetAddresses() []string {
	return c.Addresses
}
