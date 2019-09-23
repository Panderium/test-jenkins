package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.AddCommand(frameworkCmd, languageCmd, bddCmd)
}

var addCmd = &cobra.Command {
	Use: "add",
	Short: "Add framework(s), BDD(s) or language(s)",
	Long: "", // TODO
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var frameworkCmd = &cobra.Command {
	Use: "framework",
	Short: "Add a new framework to your project",
	Long: "", // TODO
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var languageCmd = &cobra.Command {
	Use: "language",
	Short: "Add a new language to your project",
	Long: "", // TODO
	Run: func(cmd *cobra.Command, args []string) {
		
	},
}

var bddCmd = &cobra.Command {
	Use: "bdd",
	Short: "Add a new bdd to your project",
	Long: "", // TODO
	Run: func(cmd *cobra.Command, args []string) {
		
	},
}