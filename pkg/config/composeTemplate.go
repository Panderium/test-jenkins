package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"text/template"
	"../tool"
)

func (conf *Config) splitConfig() (Config, Config) {
	confBack := Config{ProjectName: conf.ProjectName}
	confFront := Config{ProjectName: conf.ProjectName}

	for _, service := range conf.Services {
		if service.Name == "BDD" {
			for i, bdd := range service.Values {
				if service.Link[i] == "front" {
					confFront.UpdateServices(tool.Tool{Name: "BDD", Values: []string{bdd}})
				} else {
					confBack.UpdateServices(tool.Tool{Name: "BDD", Values: []string{bdd}})
				}
			}
		} else if service.Name == "back" {
			confBack.UpdateServices(service)
		} else {
			confFront.UpdateServices(service)
		}
	}
	return confBack, confFront
}


func (conf *Config) buildFile(templateFile, file string) {
	name := path.Base(templateFile)

	envTemplate, err := template.New(name).Funcs(template.FuncMap{
		"retrieveComposeEnvVariable": func(bddTool string) string {
			data, err := ioutil.ReadFile(".templates/BDD/" + bddTool + "/docker-compose.env.yaml")
			if err != nil {
				fmt.Printf("ENV pour la BDD non trouvés")
			}
			return string(data)
		},
		"retrieveEnvVariable": func(bddTool string) string {
			data, err := ioutil.ReadFile(".templates/BDD/" + bddTool + "/.env.bdd")
			if err != nil {
				fmt.Printf("ENV pour la BDD non trouvés")
			}
			return string(data)
		},
	}).ParseFiles(templateFile)

	if err != nil {
		fmt.Printf("template %s non reconnue\n", templateFile)
		fmt.Println(err)
		os.Exit(1)
	}

	f, err := os.Create(file)
	if err != nil {
		fmt.Printf("impossible de créer le fichier %s", file)
		os.Exit(1)
	}
	defer f.Close()

	err = envTemplate.Execute(f, conf)
	if err != nil {
		fmt.Printf("impossible de générer %s à partir de %s", file, templateFile)
		fmt.Printf("\n%s", err)
		os.Exit(1)
	}

}
func (conf *Config) helperCreateComposeAndEnv() {
	for _, service := range conf.Services {
		if service.Name != "BDD" {
			for k := range service.Values {
				dir := filepath.Join(conf.ProjectName, service.Values[k])
				if _, err := os.Stat(dir); !os.IsNotExist(err) {
					conf.buildFile(".templates/compose/docker-compose.yml", dir+"/docker-compose.yml")
					conf.buildFile(".templates/compose/docker-compose.gitlab.yml", dir+"/docker-compose.gitlab.yml")
					conf.buildFile(".templates/compose/docker-compose.gitlab.prod.yml", dir+"/docker-compose.gitlab.prod.yml")
					conf.buildFile(".templates/compose/docker-compose.preprod.yml", dir+"/docker-compose.preprod.yml")
					conf.buildFile(".templates/compose/.env", dir+"/.env")
				}
			}
		}
	}
}
// CreateComposeAndEnv create docker-compose*yml and .env for the different parts of the project
func (conf *Config) CreateComposeAndEnv() {
	confBack, confFront := conf.splitConfig()
	confBack.helperCreateComposeAndEnv()
	confFront.helperCreateComposeAndEnv()
	
}

// CreateComposeProdAndEnv create the docker-compose.prod.yml and .env files at the root directory of the project
func (conf *Config) CreateComposeProdAndEnv() {
	conf.buildFile(".templates/compose/docker-compose.prod.yml", conf.ProjectName+"/docker-compose.prod.yml")
	conf.buildFile(".templates/compose/.env", conf.ProjectName+"/.env")
}
