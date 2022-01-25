/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"log"

	"github.com/chigopher/pathlib"
	"github.com/duysqubix/gonote/gonoted"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: `Create a new note and open in editor`,
	Long: `
	Create a new note with supplied name 
	or leave blank for a quick note
	`,
	Run: NewNote,
}

func init() {
	rootCmd.AddCommand(newCmd)
}

// creates directory
// returns bool if dir is created or not
func createDir(dirPath *pathlib.Path) (bool, error) {
	dirExists, err := dirPath.IsDir()

	if !dirExists || err != nil {
		// create directory
		if err := dirPath.MkdirAll(); err != nil {
			return false, errors.New(fmt.Sprintf("unable to create directory, %s", err))
		}
	}

	return true, nil
}

func NewNote(cmd *cobra.Command, args []string) {

	var fname string

	if len(args) != 0 {
		fname = args[0]
	} else {
		fname = gonoted.GetQuickNoteName()
	}

	fname += ("." + conf.FileType)

	fullNotePath := conf.DirPath + fname

	path := pathlib.NewPathAfero(fullNotePath, afero.NewOsFs())
	dirPath := path.Parent()

	if _, err := createDir(dirPath); err != nil {
		log.Fatalf("Unable to create directory, %v", err)
	}

	if err := gonoted.OpenEditorSync(cmd, &conf, fullNotePath); err != nil {
		log.Fatalf("Unable to open editor, %v", err)
	}
}
