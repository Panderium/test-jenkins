package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"../pkg/config"
	"../pkg/tool"

	"github.com/spf13/cobra"
)

// Tools flag for adding tool
var Tools string

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add [BDD/front/back] [path to your project]",
	Short: "Add BDD and language or framework to your backend or frontend",
	Long:  "",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(args[1]); os.IsNotExist(err) {
			fmt.Println("Chemin vers le projet invalide")
			os.Exit(1)
		}
		if args[0] != "BDD" && args[0] != "front" && args[0] != "back" {
			fmt.Println("Nom de service invalide, utiliser BDD, front ou back")
			os.Exit(1)
		}

		config.CloneTemplates()
		defer os.RemoveAll(".templates")

		conf := config.LoadConfigFile(args[1])
		newTool := tool.Tool{Name: args[0]}

		idxBDD := conf.SearchServiceIndex("BDD")
		idxBack := conf.SearchServiceIndex("back")
		idxFront := conf.SearchServiceIndex("front")

		newTool.Select()

		conf.UpdateServices(newTool)

		if idxBDD != -1 {
			if idxBack != -1 {
				conf.Services[idxBDD].LinkWith(&conf.Services[idxBack])
			}
			if idxFront != -1 {
				conf.Services[idxBDD].LinkWith(&conf.Services[idxFront])
			}
		}

		// retrieve files and put it in  back and front folders
		conf.RetrieveFiles()

		// build .config.yaml
		yamlConf := conf.BuildConfigFile()
		ioutil.WriteFile(conf.ProjectName+"/.config.yaml", yamlConf, 0644)

		// create docker-compose docker-compose.gitlab docker-compose.preprod with .env
		conf.CreateComposeAndEnv()

		// create docker-compose.prod with .env
		conf.CreateComposeProdAndEnv()

		// git init
		conf.GitInit()

	},
}

