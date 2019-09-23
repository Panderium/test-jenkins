package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
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
		fmt.Println("Création d'un nouveau projet nommé %s", args[0])
		// create root directory with children
		// les bases de données
		// le backend
		// le frontend
		// Setup .env file for env variables
		// create docker-compose.prod.yaml (bdd, back, front) 

	},
}