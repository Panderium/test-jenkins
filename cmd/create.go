package cmd

import (
	"fmt"
	"os"
	"io/ioutil"

	"../pkg/tool"
	"../pkg/config"

	git "gopkg.in/src-d/go-git.v4"
	"github.com/spf13/cobra"
)


func init() {
	rootCmd.AddCommand(createCmd)

	_, err := git.PlainClone("../.templates", false, &git.CloneOptions{
		URL: "http://10.1.38.31/afougerouse/templates.git",
		Progress: os.Stdout,
	})
	if err != nil {
		fmt.Errorf("Impossible de récupérer les templates")
		os.Exit(1)
	}
}

var createCmd = &cobra.Command{
	Use: "create",
	Short: "Create an empty project with according config files",
	Long: "",// TODO
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("Entrer le nom du projet en argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		
		defer os.RemoveAll("../.templates")
		
		projectRoot := args[0]
		fmt.Println("Création d'un nouveau projet nommé ", projectRoot)
		/// SETUP ALL YAML FILES WITH ACCORDING TODO .config.yaml ///
		// create root directory and children (?usefull to have directory path?), init value
		bdd := tool.Tool{Name: "bdd", Values: nil, Link: nil}
		back := tool.Tool{Name: "back", Values: nil, Link: nil}
		front := tool.Tool{Name: "front", Values: nil, Link: nil}

		conf:= config.Config{}

		os.MkdirAll(projectRoot + "/back", 0777)
		os.MkdirAll(projectRoot + "/front", 0777)

		// update config file with project name
		conf.UpdateProjectName(projectRoot)
		
		// Select tools
		bdd.Select()
		back.Select()
		front.Select()

		// Link bdd(s) with front and/or back
		bdd.LinkWith(&back)
		bdd.LinkWith(&front)

		// update config files with tools
		conf.UpdateToolConfig(bdd)
		conf.UpdateToolConfig(back)
		conf.UpdateToolConfig(front)

		// build .config.yaml
		yamlConf := conf.BuildConfigFile()
		ioutil.WriteFile(projectRoot + "/.config.yaml", yamlConf, 0644)

		
		// Setup .env file for env variables
		// create docker-compose.prod.yaml (bdd, back, front if existe) 
	},
}