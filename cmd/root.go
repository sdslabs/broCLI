package cmd

import (
	"os"

	"github.com/sdslabs/broCLI/logger"

	"github.com/spf13/cobra"
)

var (
	log = logger.NewLogger(os.Stdout)

	config string // Flag

	gamePath string // Set when running any command if Game path exists

	rootCmd = &cobra.Command{
		Use:              "bro",
		Short:            "A command line tool to streamline your game development process with Rubeus.",
		Long:             `BroCLI is a tool meant to make your development process easier with commands to create, build and run your project maintaining the proper architecture for Rubeus game.`,
		PersistentPreRun: persistentPreRun,
		Run: func(cmd *cobra.Command, args []string) {
			log.Warn("No arguments found... Here's help:")
			cmd.Help()
		},
	}
)

// Execute commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	// Flags
	rootCmd.Flags().StringVarP(&config, "config", "c", "", "set Game dir config")

	// Commands
	rootCmd.AddCommand(createGame)
	rootCmd.AddCommand(createLevel)
	rootCmd.AddCommand(createObject)
}
