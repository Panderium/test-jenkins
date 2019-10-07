package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
)

// GitInit perform a git init in the folders of the project and add a remote according to the will of the holy user
func (c *Config) GitInit() {
	for _, service := range c.Services {
		if service.Name != "BDD" {
			for k := range service.Values {
				dir := filepath.Join(c.ProjectName, service.Values[k])
				if _, err := os.Stat(dir + "/.git"); !os.IsNotExist(err) {
					continue
				}
				repo, err := git.PlainInit(dir, false)
				if err != nil {
					fmt.Println(err)
				}

				value := ""
				prompt := &survey.Select{
					Message: "Il y a-t-il un repo GitLab pour le " + service.Name + "/" + service.Values[k] + " ?",
					Options: []string{"Oui", "Non"},
				}
				survey.AskOne(prompt, &value)

				if value == "Oui" {
					var urlGitlab string
					prompt := &survey.Input{
						Message: "Entrer l'URL http GitLab",
					}
					survey.AskOne(prompt, &urlGitlab)

					_, err = repo.CreateRemote(&config.RemoteConfig{
						Name: "origin",
						URLs: []string{urlGitlab},
					})
				}
			}
		}
	}
}

// CloneTemplates clone the repo from Gitlab which containes all the needed templates and other files like Dockerfiles.
func CloneTemplates() {
	_, err := git.PlainClone(".templates", false, &git.CloneOptions{
		URL:      "http://10.1.38.31/afougerouse/templates.git",
		Progress: os.Stdout,
	})
	if err != nil {
		fmt.Errorf("Impossible de récupérer les templates")
		os.Exit(1)
	}
}
