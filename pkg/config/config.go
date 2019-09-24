package config

import (
	"fmt"
	"os"

	"../tool"

	"github.com/AlecAivazis/survey"
	"gopkg.in/yaml.v2"
)

// Config TODO
type Config struct {
	Bdd   []string
	Back  []string
	Front []string
	Link  map[string]string
}

// BddLinkWith TODO
func (c *Config) BddLinkWith() {
	var linkedTo string
	c.Link = make(map[string]string)

	for _, bdd := range c.Bdd {
		prompt := &survey.Select{
			Message: "Avec quelle partie du projet la base de données " + bdd + " est-elle reliée ?",
			Options: []string{"back", "front"},
		}

		survey.AskOne(prompt, &linkedTo)
		c.Link[bdd] = linkedTo
	}
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
func (c *Config) BuildConfigFile() []byte {
	yamlFile, err := yaml.Marshal(&c)
	if err != nil {
		fmt.Printf("Impossible de construire le fichier de config")
		os.Exit(1)
	}
	return yamlFile
}
