package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"../pkg/config"
	"../pkg/tool"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
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
		config.CloneTemplates()

		defer os.RemoveAll(".templates")

		projectRoot := args[0]
		fmt.Println("Création d'un nouveau projet nommé ", projectRoot)
		/// SETUP ALL YAML FILES WITH ACCORDING TODO .config.yaml ///
		// create root directory and children (?usefull to have directory path?), init value
		bdd := tool.Tool{Name: "BDD", Values: nil, Link: nil}
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

		// retrieve files and put it in  back and front folders
		conf.RetrieveFiles()

		// build .config.yaml
		yamlConf := conf.BuildConfigFile()
		ioutil.WriteFile(projectRoot+"/.config.yaml", yamlConf, 0644)

		// create docker-compose docker-compose.gitlab docker-compose.preprod with .env
		conf.CreateComposeAndEnv()

		// create docker-compose.prod with .env
		conf.CreateComposeProdAndEnv()

		// git init
		conf.GitInit()

	},
}
