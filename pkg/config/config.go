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
	Services    []tool.Tool
}

func addCIFiles(dest string) error {
	gitlabCIFile := ".templates/CI/.gitlab-ci.yml"
	sonarQubeFile := ".templates/CI/sonar-scanner.sh"

	if err :=utils.CopyFile(gitlabCIFile, dest + "/.gitlab-ci.yml"); err != nil {
		return err
	}
	if err := utils.CopyFile(sonarQubeFile, dest + "/sonar-scanner.sh"); err != nil {
		return err
	}

	return nil
}

// RetrieveFiles TODO
func (c *Config) RetrieveFiles() {
	var src string
	var dest string

	for _, service := range c.Services {
		if service.Name != "BDD" {
			src = ".templates/" + service.Name + "/" + service.Values[0]
			dest = c.ProjectName + "/" + service.Name
			err := utils.CopyDir(src, dest)
			if err != nil {
				fmt.Printf("imposible de récupérer les éléments pour construire le %s\n", service.Name)
			}
			err = addCIFiles(dest)
			if err != nil {
				fmt.Printf("imposible de récupérer les fichiers de CI/CD pour le %s\n", service.Name)
				fmt.Println(err)
			}
		}
	}
}

// UpdateProjectName TODO
func (c *Config) UpdateProjectName(name string) {
	c.ProjectName = name
}

// UpdateServices TODO
func (c *Config) UpdateServices(t tool.Tool) {
	if t.Values != nil && t.Values[0] != "aucune" {
		c.Services = append(c.Services, t)
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
