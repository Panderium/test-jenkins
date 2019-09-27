package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"text/template"

	"../tool"
)

func (conf *Config) splitConfig() (Config, Config) {
	confBack := Config{ProjectName: conf.ProjectName}
	confFront := Config{ProjectName: conf.ProjectName}

	for _, service := range conf.Services {
		if service.Name == "BDD" {
			for i, bdd := range service.Values {
				if service.Link[i] == "Front" {
					confFront.UpdateServices(tool.Tool{Name: "BDD", Values: []string{bdd}})
				} else {
					confBack.UpdateServices(tool.Tool{Name: "BDD", Values: []string{bdd}})
				}
			}
		} else if service.Name == "Back" {
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

// CreateComposeAndEnv TODO
func (conf *Config) CreateComposeAndEnv() {
	confBack, confFront := conf.splitConfig()
	if _, err := os.Stat(confBack.ProjectName + "/Back"); !os.IsNotExist(err) {
		confBack.buildFile("templates/docker-compose.yml", confBack.ProjectName+"/Back/docker-compose.yml")
		confBack.buildFile("templates/docker-compose.gitlab.yml", confBack.ProjectName+"/Back/docker-compose.gitlab.yml")
		confBack.buildFile("templates/docker-compose.preprod.yml", confBack.ProjectName+"/Back/docker-compose.preprod.yml")
		confBack.buildFile("templates/.env", confBack.ProjectName+"/Back/.env")
	}
	if _, err := os.Stat(confBack.ProjectName + "/Front"); !os.IsNotExist(err) {
		confFront.buildFile("templates/docker-compose.yml", confFront.ProjectName+"/Front/docker-compose.yml")
		confFront.buildFile("templates/docker-compose.gitlab.yml", confFront.ProjectName+"/Front/docker-compose.gitlab.yml")
		confFront.buildFile("templates/docker-compose.preprod.yml", confFront.ProjectName+"/Front/docker-compose.preprod.yml")
		confFront.buildFile("templates/.env", confFront.ProjectName+"/Front/.env")
	}
}

// CreateComposeProdAndEnv TODO
func (conf *Config) CreateComposeProdAndEnv() {
	conf.buildFile("templates/docker-compose.prod.yml", conf.ProjectName+"/docker-compose.prod.yml")
	conf.buildFile("templates/.env", conf.ProjectName+"/.env")
}
