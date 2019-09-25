package config

import (
	"fmt"
	"os"

	"../tool"

	"gopkg.in/yaml.v2"
)

// Config TODO
type Config struct {
	ProjectName string
	Services []tool.Tool
}

// UpdateProjectName TODO
func (c *Config) UpdateProjectName(name string) {
	c.ProjectName = name
}

// UpdateToolConfig TODO
func (c *Config) UpdateToolConfig(t tool.Tool) {
	c.Services = append(c.Services, t)
}

// BuildConfigFile TODO
func (c *Config) BuildConfigFile() []byte {
	yamlFile, err := yaml.Marshal(&c)
	if err != nil {
		fmt.Printf("Impossible de construire le fichier de config")
		os.Exit(1)
	}
	return yamlFile
}
