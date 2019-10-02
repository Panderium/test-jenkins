package config

import (
	"fmt"

	"github.com/AlecAivazis/survey"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
)

// GitInit perform a git init in the folders of the project and add a remote according to the will of the holy user
func (c *Config) GitInit() {
	for _, service := range c.Services {
		if service.Name != "BDD" {
			repo, err := git.PlainInit(c.ProjectName+"/"+service.Name, false)
			if err != nil {
				fmt.Println(err)
			}

			value := ""
			prompt := &survey.Select{
				Message: "Il y a-t-il un repo GitLab pour le " + service.Name + " ?",
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
