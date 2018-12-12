package cmd

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sdslabs/broCLI/utils"
	"github.com/spf13/cobra"
)

var createGame = &cobra.Command{
	Use:   "creategame [title]",
	Short: "creates new Rubeus game",
	Long:  "Creates new Rubeus game.",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("atleast one argument required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		path := filepath.Join(gamePath, name)
		// Check if name is valid
		if !utils.IsValidName(name) {
			log.Fatalf("'%s' is not valid.%sGame titles should only include alphanumeric characters or underscore.", name, utils.NLn)
		}
		// Check if game with the same name already exists
		if utils.DoesExist(path) && utils.IsDir(path) {
			log.Fatalf("'%s' Game already exists! [%s]", name, path)
		}
		var err error
		// Check if 'user_init.h' and 'user_init.cpp' exist, create if not
		userInitH := filepath.Join(gamePath, "user_init.h")
		if !(utils.DoesExist(userInitH) && utils.IsFile(userInitH)) {
			log.Info("Creating 'user_init.h'...")
			err = ioutil.WriteFile(userInitH, []byte(utils.UserInitH), 0644)
			if err != nil {
				log.Error("Failed to create 'user_init.h'")
			}
		}
		userInitCpp := filepath.Join(gamePath, "user_init.cpp")
		if !(utils.DoesExist(userInitCpp) && utils.IsFile(userInitCpp)) {
			log.Info("Creating 'user_init.cpp'...")
			err = ioutil.WriteFile(userInitCpp, []byte(utils.UserInitCpp), 0644)
			if err != nil {
				log.Error("Failed to create 'user_init.cpp'")
			}
		}
		// Create a folder named name and name/engine_files
		// Path error should not exist (already verified)
		engFiles := filepath.Join(path, "engine_files")
		err = os.MkdirAll(engFiles, 0755)
		if err != nil {
			log.Fatalf("Could not create %s! Path does not exist.", path)
		}
		log.Logf("New game '%s' created at %s", name, path)
	},
}
