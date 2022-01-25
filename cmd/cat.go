/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/duysqubix/gonote/gonoted"
	"github.com/spf13/cobra"
)

// catCmd represents the cat command
var catCmd = &cobra.Command{
	Use:   "cat",
	Short: "Displays note",
	Run: func(cmd *cobra.Command, args []string) {
		var fname string
		if len(args) == 0 {
			fname = gonoted.GetQuickNoteName()
		} else {
			fname = args[0]
		}

		notePath := fmt.Sprintf("%s%s.%s", conf.DirPath, fname, conf.FileType)
		notePObj := gonoted.MakePathObj(notePath)

		if exists, _ := notePObj.Exists(); !exists {
			cmd.Printf("no such note")
			os.Exit(1)
		}

		if isFile, _ := notePObj.IsFile(); !isFile {
			cmd.Printf("Can not read a directory")
			os.Exit(1)
		}

		contents, _ := notePObj.ReadFile()

		fmt.Println(string(contents))
	},
}

func init() {
	rootCmd.AddCommand(catCmd)
}
