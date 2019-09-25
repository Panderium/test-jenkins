package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"../pkg/config"
	"../pkg/tool"

	"github.com/spf13/cobra"
	git "gopkg.in/src-d/go-git.v4"
)

func init() {
	rootCmd.AddCommand(createCmd)

	_, err := git.PlainClone(".templates", false, &git.CloneOptions{
		URL:      "http://10.1.38.31/afougerouse/templates.git",
		Progress: os.Stdout,
	})
	if err != nil {
		fmt.Errorf("Impossible de récupérer les templates")
		os.Exit(1)
	}
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an empty project with according config files",
	Long:  "", // TODO
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("Entrer le nom du projet en argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		defer os.RemoveAll(".templates")

		projectRoot := args[0]
		fmt.Println("Création d'un nouveau projet nommé ", projectRoot)
		/// SETUP ALL YAML FILES WITH ACCORDING TODO .config.yaml ///
		// create root directory and children (?usefull to have directory path?), init value
		bdd := tool.Tool{Name: "bdd", Values: nil, Link: nil}
		back := tool.Tool{Name: "back", Values: nil, Link: nil}
		front := tool.Tool{Name: "front", Values: nil, Link: nil}

		conf := config.Config{}

		// update config file with project name
		conf.UpdateProjectName(projectRoot)

		// Select tools
		bdd.Select()
		back.Select()
		front.Select()

		// Link bdd(s) with front and/or back
		bdd.LinkWith(&back)
		bdd.LinkWith(&front)

		// update config files with tools(services)
		conf.UpdateServices(bdd)
		conf.UpdateServices(back)
		conf.UpdateServices(front)

		// retrieve files and put it in tmp/bdd, back and front folders
		conf.RetrieveBdd()
		conf.Retrieve("back")
		conf.Retrieve("front")

		// build .config.yaml
		yamlConf := conf.BuildConfigFile()
		ioutil.WriteFile(projectRoot+"/.config.yaml", yamlConf, 0644)

		// create docker-compose docker-compose.gitlab docker-compose.preprod with .env
		conf.CreateComposeAndEnv()

		// create docker-compose.prod with .env
		conf.CreateComposeProdAndEnv()

		// clean up

		// git init

	},
}
