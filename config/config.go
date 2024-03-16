package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type config struct {
	File  File   `yaml:"file"`
	Pizza Pizza  `yaml:"pizza"`
	Rooms []Room `yaml:"rooms"`
}

type File struct {
	Name string `yaml:"name"`
}

type Pizza struct {
	SlicesPerPizza    int `yaml:"slicesPerPizza"`
	ExtraCheeseSlices int `yaml:"extraCheeseSlices"`
}

type Room struct {
	Teacher string `yaml:"teacher"`
	Room    string `yaml:"room"`
	Class   string `yaml:"class"`
	Code    string `yaml:"code"`
}

func ReadConfig(configPath string) (config, error) {
	f, err := os.Open(configPath)
	if errors.Is(err, os.ErrNotExist) {
		return config{}, fmt.Errorf("no such file %s", configPath)
	}
	if err != nil {
		return config{}, fmt.Errorf("failed to open config file %s: %w", configPath, err)
	}
	defer f.Close()

	var conf config
	err = yaml.NewDecoder(f).Decode(&conf)
	if err != nil {
		return conf, fmt.Errorf("failed to read config file: %w", err)
	}

	return conf, nil
}
