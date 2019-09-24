package config

import (
	"os"
	"fmt"
	//"io/ioutil"
	"../tool"
	"gopkg.in/yaml.v2"
)

// Config TODO
type Config struct {
	Bdd []string
	Back []string
	Front []string
}

// UpdateToolConfig TODO
func (c *Config) UpdateToolConfig(t *tool.Tool) {
	switch t.Name {
	case "BDD":
		c.Bdd = append(c.Bdd, t.Values...)
	case "Back":
		c.Back = append(c.Back, t.Values...)
	case "Front":
		c.Front = append(c.Front, t.Values...)
	default:
		fmt.Println("%s non pris en charge", t.Name)
	}
}

// BuildConfigFile TODO
func (c *Config) BuildConfigFile() {
	yamlFile, err := yaml.Marshal(&c)
	if err != nil {
		fmt.Printf("Impossible de construire le fichier de config")
		os.Exit(1)
	}
	fmt.Printf(string(yamlFile))
}