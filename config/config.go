package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)


func Parse(config *os.File) (*Config, error) {
	c := &Config{}
	d := yaml.NewDecoder(config)
	err := d.Decode(c)
	if err != nil {
		return nil, err
	}

	err = c.Validate()
	if err != nil {
		return nil, fmt.Errorf("validating top-level config: %q", err)
	}

	err = c.Branch.Validate()
	if err != nil {
		return nil, fmt.Errorf("validating branch config: %q", err)
	}

	err = c.Commit.Validate()
	if err != nil {
		return nil, fmt.Errorf("validating commit config: %q", err)
	}

	if c.Push != nil {
		err = c.Push.Validate()
		if err != nil {
			return nil, fmt.Errorf("validating push config: %q", err)
		}
	}

	if c.PR != nil {
		err = c.PR.Validate()
		if err != nil {
			return nil, fmt.Errorf("validating PR config: %q", err)
		}
	}
	return c, nil
}

// struct and fields need to be public for marshalling/unmarshalling
type Config struct {
	Branch BranchConfig `yaml:"branch"`
	Commit CommitConfig `yaml:"commit"`
	Push *PushConfig `yaml:"push"`
	PR *PRConfig `yaml:"pr"`
}

func (c *Config) Validate() error {
	return nil
}


type PushConfig struct {
}

func (c *PushConfig) Validate() error {
	return nil
}




