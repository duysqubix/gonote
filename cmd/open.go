/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/chigopher/pathlib"
	"github.com/duysqubix/gonote/gonoted"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Opens an existing note by name",

	Run: runOpen,
}

func init() {
	rootCmd.AddCommand(openCmd)
}

func runOpen(cmd *cobra.Command, args []string) {
	var fname string
	if len(args) == 0 {
		// default to quicknote name for today
		fname = gonoted.GetQuickNoteName()
	} else {
		fname = args[0]
	}

	notePath := fmt.Sprintf("%s%s.%s", conf.DirPath, fname, conf.FileType)

	notePObj := pathlib.NewPathAfero(notePath, afero.NewOsFs())

	if exists, _ := notePObj.Exists(); !exists {
		cmd.Printf("note: %s not found, create with `gonote new %s`\n", fname+"."+conf.FileType, fname)
		os.Exit(1)
	}

	if isFile, _ := notePObj.IsFile(); !isFile {
		log.Fatalf("targted object is not a valid file. given, %v", notePath)
	}

	if err := gonoted.OpenEditorSync(cmd, &conf, notePath); err != nil {
		log.Fatal(err)
	}
}
