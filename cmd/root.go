package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "brocli",
	Short: "A command line tool to streamline your game development process with Rubeus.",
	Long:  `BroCLI is a tool meant to make your development process easier with commands to create, build and run your project maintaining the proper architecture for Rubeus game.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to broCLI.")
	},
}

// Execute commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
