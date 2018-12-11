package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sdslabs/broCLI/utils"
	"github.com/spf13/cobra"
)

func persistentPreRun(cmd *cobra.Command, args []string) {
	// check if config exists, if not ask for one
	if config != "" {
		if !utils.IsAbsolute(config) {
			fmt.Printf("config path '%s' is not absolute\n", config)
			os.Exit(1)
		}
		if !utils.DoesExist(config) {
			fmt.Printf("config path '%s' does not exist\n", config)
			os.Exit(1)
		}
		if !utils.IsDir(config) {
			fmt.Printf("config path '%s' is not a directory\n", config)
			os.Exit(1)
		}
		gamePath = config
		// set config file
		err := ioutil.WriteFile(utils.BroConfPath, []byte(gamePath), 0644)
		if err != nil {
			fmt.Printf("error while reading config file\n")
			os.Exit(1)
		}
		fmt.Printf("config written successfully\n")
		os.Exit(0)
	}
	if utils.DoesExist(utils.BroConfPath) {
		file, err := ioutil.ReadFile(utils.BroConfPath)
		if err != nil {
			fmt.Printf("error while reading config file\n")
			os.Exit(1)
		}
		gamePath = string(file)
		return
	}
	fmt.Printf("config does not exist\n")
	os.Exit(1)
}
