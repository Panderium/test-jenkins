package cmd

import (
	"fmt"
	"os"

	slct "../pkg"
	Tool "../pkg"

	git "github.com/src-d/go-git"
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
		/// SETUP ALL YAML FILES WITH ACCORDING TODO .config.yaml ///
		// create root directory and children (?usefull to have directory path?), init value
		bdd := Tool{"bdd", nil}
		back := Tool{"backend", nil}
		front := Tool{"frontend", nil}

		os.MkdirAll(args[0] + "/back", 0777)
		os.MkdirAll(args[0] + "/front", 0777)

		// Select tools
		// les bases de données
		bdd.Select()
		// le backend
		back.Select()
		// le frontend
		front.Select()
		// where to link bdd ?

		// Setup .env file for env variables
		// create docker-compose.prod.yaml (bdd, back, front if existe) 
	},
}