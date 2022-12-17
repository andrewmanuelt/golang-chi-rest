package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func ConfigField(configuration string, field string) string {
	loader := loader()

	maps := loader[configuration].(map[string]interface{})

	return maps[field].(string)
}

func loader() map[string]interface{} {
	config, err := ioutil.ReadFile("config.yaml")

	if err != nil {
		panic(err)
	}

	data := make(map[string]interface{})

	err = yaml.Unmarshal(config, &data)

	if err != nil {
		panic(err)
	}

	return data
}
