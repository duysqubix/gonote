/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	gnote "github.com/duysqubix/gonote/gonoted"
	"github.com/spf13/cobra"
)

const (
	usage = `
	gonote new <note-name>
	gonote ls
	`
	short         = `Simple note taking system within the terminal using your favorite terminal editor`
	long          = `THis is a long description`
	defaultEditor = "nano"
)

var conf gnote.Config

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gonote",
	Short: short,
	Long:  long,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	userHome := os.Getenv("HOME")
	if userHome == "" {
		fmt.Fprintln(os.Stderr, "$HOME variable is not set")
		os.Exit(1)
	}

	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = defaultEditor
	}
	rootCmd.PersistentFlags().SortFlags = false
	rootCmd.PersistentFlags().StringVarP(&conf.Editor, "editor", "e", editor, "Selection of editor")
	rootCmd.PersistentFlags().StringVarP(&conf.DirPath, "notes-directory", "d", userHome+"/.notes/", "Sets directory of main notes")
	rootCmd.PersistentFlags().StringVar(&conf.FileType, "default-file-type", "md", "Sets default file type for new notes")
}
