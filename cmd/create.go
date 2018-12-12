package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/sdslabs/broCLI/utils"
)

// for creating object or level
func create(name, game, what, wH, wCpp string) {
	// Check if variable name is valid
	if !utils.IsValidVarName(name) {
		log.Fatalf("'%s' is not a valid variable name!%sVariable names can have alphanumeric characters and underscore only.%sFirst letter should not be a number.", name, utils.NLn, utils.NLn)
	}
	// Check if engine_files dir exists
	engFiles := filepath.Join(gamePath, game, "engine_files")
	if !(utils.DoesExist(engFiles) && utils.IsDir(engFiles)) {
		log.Fatalf("Path '%s' does not exist. Make sure it does.%sPossibly the game does not exist. Create using 'bro creategame %s'", engFiles, utils.NLn, game)
	}
	// Check if it already exists (both .h and .cpp)
	pathH := filepath.Join(engFiles, fmt.Sprintf("%s.%s.h", what, name))
	if utils.DoesExist(pathH) && utils.IsFile(pathH) {
		log.Fatalf("%s '%s' already exists! [%s]", strings.Title(what), name, pathH)
	}
	pathCpp := filepath.Join(engFiles, fmt.Sprintf("%s.%s.cpp", what, name))
	if utils.DoesExist(pathCpp) && utils.IsFile(pathCpp) {
		log.Fatalf("%s '%s' already exists! [%s]", strings.Title(what), name, pathCpp)
	}
	var err error
	// Create what.name.h and what.name.cpp
	err = ioutil.WriteFile(pathH, []byte(wH), 0644)
	if err != nil {
		log.Errorf("Failed to create '%s.%s.h'", what, name)
	}
	err = ioutil.WriteFile(pathCpp, []byte(wCpp), 0644)
	if err != nil {
		log.Errorf("Failed to create '%s.%s.cpp'", what, name)
	}
	log.Logf("New %s '%s' created!", what, name)
}

var createLevel = &cobra.Command{
	Use:   "createlevel [game] [level name]",
	Short: "creates new Rubeus game level",
	Long:  "Creates new Rubeus game level.",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("atleast two arguments required, game and level name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		name := args[1]
		game := args[0]
		what := "level"
		hname := fmt.Sprintf("%s.%s", what, name)
		// Change when files are updated
		levelH := fmt.Sprintf(utils.LevelH, name, name, name)
		levelCpp := fmt.Sprintf(utils.LevelCpp, hname, name, name)

		create(name, game, what, levelH, levelCpp)
	},
}

var createObject = &cobra.Command{
	Use:   "createobject [game] [object name]",
	Short: "creates new Rubeus game object",
	Long:  "Creates new Rubeus game object.",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("atleast two arguments required, game and object name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		name := args[1]
		game := args[0]
		what := "object"
		hname := fmt.Sprintf("%s.%s", what, name)
		// Change when files are updated
		objectH := fmt.Sprintf(utils.ObjectH, name, name, name)
		objectCpp := fmt.Sprintf(utils.ObjectCpp, hname, name, name, name, name)

		create(name, game, what, objectH, objectCpp)
	},
}
