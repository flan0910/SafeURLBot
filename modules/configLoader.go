package modules

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type config struct {
	ApiKey string `yaml:"ApiKey"`
	Token string `yaml:"Token"`
}

func ConfigLoader() config {
	config := config{}
	b, err := os.ReadFile("./config.yaml")
	if err != nil {
		fmt.Println(err)
	}
	yaml.Unmarshal(b, &config)
	
	return config
}