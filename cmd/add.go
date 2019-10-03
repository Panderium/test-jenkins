package cmd

import (
	"os"
	"fmt"
	
	"../pkg/tool"
	"../pkg/config"

	"github.com/spf13/cobra"
)

// Tools flag for adding tool
var Tools string

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command {
	Use: "add [bbd/front/back] [path to your to project]",
	Short: "Add BDD and language or framework to your backend or frontend",
	Long: "",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(args[1]); os.IsNotExist(err) {
			fmt.Println("Chemin vers le projet invalide")
			os.Exit(1)
		}
		
	},
}

var bbdCmd = &cobra.Command {
	Use: "bdd",
	Short: "Add a new bdd to your project",
	Long: "",
	Run: func(cmd *cobra.Command, args []string) {
		conf := config.LoadConfigFile(args[0])
		conf.Services[].select()
		
	},
}

var backCmd = &cobra.Command {
	Use: "back",
	Short: "Add a new back to your project",
	Long: "",
	Run: func(cmd *cobra.Command, args []string) {
		conf := config.LoadConfigFile(args[0])
		conf.Services[].select()
	},
}

var frontCmd = &cobra.Command {
	Use: "front",
	Short: "Add a new front to your project",
	Long: "",
	Run: func(cmd *cobra.Command, args []string) {
		conf := config.LoadConfigFile(args[0])
		conf.Services[].select()
	},
}