package cmd

import (
	"io/ioutil"
	"os"

	"github.com/sdslabs/broCLI/utils"
	"github.com/spf13/cobra"
)

func persistentPreRun(cmd *cobra.Command, args []string) {
	// check if config exists, if not ask for one
	if config != "" {
		checkPath(config)
		gamePath = config
		// set config file
		err := ioutil.WriteFile(utils.BroConfPath, []byte(gamePath), 0644)
		if err != nil {
			log.Fatal("Error while writing config file")
		}
		log.Log("Config written successfully")
		os.Exit(0)
	}
	// While running any other command
	if utils.DoesExist(utils.BroConfPath) {
		file, err := ioutil.ReadFile(utils.BroConfPath)
		if err != nil {
			log.Fatal("Error while reading config file")
		}
		gamePath = string(file)
		checkPath(gamePath)
		return
	}
	log.Fatalf("Game path does not exist.%sSet it using '%s --config /path/to/Game'", utils.NLn, cmd.Name())
}

func checkPath(config string) {
	if !utils.IsAbsolute(config) {
		log.Fatalf("Config path '%s' is not absolute", config)
	}
	if !utils.DoesExist(config) {
		log.Fatalf("Config path '%s' does not exist", config)
	}
	if !utils.IsDir(config) {
		log.Fatalf("Config path '%s' is not a directory", config)
	}
}
