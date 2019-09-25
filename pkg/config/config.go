package config

import (
	"fmt"
	"os"

	"../tool"
	"../utils"

	"gopkg.in/yaml.v2"
)

// Config TODO
type Config struct {
	ProjectName string
	Services []tool.Tool
}

// RetrieveBdd TODO
func (c *Config) RetrieveBdd() {
	for _, bdd := range c.Services[0].Values {
		c.Retrieve(bdd)
	}
}

// Retrieve TODO
func (c *Config) Retrieve(tool string) {
	var src string
	var dest string

	switch tool {
	case "back":
		src = ".templates/Back/" + c.Services[1].Values[0]
		dest =  c.ProjectName + "/Back"
	case "front":
		src = ".templates/Front/" + c.Services[2].Values[0]
		dest =  c.ProjectName + "/Front"
	default:
		src = ".templates/BDD/" + tool
		dest =  c.ProjectName + "/tmp/.bdd"
	}
	fmt.Println(src)
	fmt.Println(dest)
	err := utils.CopyDir(src, dest)
	if err != nil {
		fmt.Printf("imposible de récupérer les éléments pour construire le %s", tool)
	}
}

// UpdateProjectName TODO
func (c *Config) UpdateProjectName(name string) {
	c.ProjectName = name
}

// UpdateServices TODO
func (c *Config) UpdateServices(t tool.Tool) {
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
