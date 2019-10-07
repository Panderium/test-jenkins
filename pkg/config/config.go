package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"../tool"
	"../utils"

	"gopkg.in/yaml.v2"
)

// Config structure of the .conf.yaml files. It owns all the needs to construct all docker-compose files
type Config struct {
	ProjectName string
	Services    []tool.Tool
}

func addCIFiles(dest string) error {
	gitlabCIFile := ".templates/CI/.gitlab-ci.yml"
	sonarQubeFile := ".templates/CI/sonar-scanner.sh"

	if err := utils.CopyFile(gitlabCIFile, dest+"/.gitlab-ci.yml"); err != nil {
		return err
	}
	if err := utils.CopyFile(sonarQubeFile, dest+"/sonar-scanner.sh"); err != nil {
		return err
	}

	return nil
}

// RetrieveFiles retrieve all the files needed to build and configure technologies chosen by the user
// e.g. Dockerfiles, create app, etc...
func (c *Config) RetrieveFiles() {
	var src string
	var dest string

	for _, service := range c.Services {
		if service.Name != "BDD" {
			for k := range service.Values {
				src = filepath.Join(".templates", service.Name, service.Values[k])
				dest = filepath.Join(c.ProjectName, service.Name, service.Values[k])
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
}

// UpdateProjectName change the variable ProjectName in c with name
func (c *Config) UpdateProjectName(name string) {
	c.ProjectName = name
}

// UpdateServices add tool t to the array Services of c
func (c *Config) UpdateServices(t tool.Tool) {
	if t.Values != nil && t.Values[0] != "aucune" {
		c.Services = append(c.Services, t)
	}
}

// BuildConfigFile build the files .conf.yaml thanks to the Config c.
func (c *Config) BuildConfigFile() []byte {
	yamlFile, err := yaml.Marshal(&c)
	if err != nil {
		fmt.Printf("Impossible de construire le fichier de config")
		os.Exit(1)
	}
	return yamlFile
}

// LoadConfigFile load and put the config files of the project in path
func LoadConfigFile(path string) Config {
	config := Config{}
	yamlFile, err := ioutil.ReadFile(path + "/.conf.yaml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return config
}
