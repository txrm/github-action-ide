package actions

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

// WIP, trimmed spec subset

type Input struct {
	Description string      `yaml:"description"`
	Default     interface{} `yaml:"default,omitempty"`
}

type Output struct {
	Description string `yaml:"description"`
}

type Runs struct {
	Using string `yaml:"using"`
	Main  string `yaml:"main"`
	Post  string `yaml:"post,omitempty"`
}

type GithubAction struct {
	Name        string            `yaml:"name"`
	Description string            `yaml:"description"`
	Inputs      map[string]Input  `yaml:"inputs"`
	Outputs     map[string]Output `yaml:"outputs"`
	Runs        Runs              `yaml:"runs"`
}

func ParseActionYML(yamlContent string) (*GithubAction, error) {
	var action GithubAction
	err := yaml.Unmarshal([]byte(yamlContent), &action)
	if err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %v", err)
	}
	return &action, nil
}
