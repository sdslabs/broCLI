package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	config string // Flag

	gamePath string

	rootCmd = &cobra.Command{
		Use:              "bro",
		Short:            "A command line tool to streamline your game development process with Rubeus.",
		Long:             `BroCLI is a tool meant to make your development process easier with commands to create, build and run your project maintaining the proper architecture for Rubeus game.`,
		PersistentPreRun: persistentPreRun,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

// Execute commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&config, "config", "c", "", "/absolute/path/rubeus/Game")
}
