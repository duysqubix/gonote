package main

import (
	"fmt"
	"os"

	"github.com/duysqubix/gonote/gonoted"
	"github.com/spf13/cobra"
)

var conf gonoted.Config

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

var rootCmd = &cobra.Command{
	Use:     usage,
	Example: "Example",
	Short:   "Short",
	Version: "0.1",
	Run:     rootRun,
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
	rootCmd.Flags().SortFlags = false
	rootCmd.Flags().StringVarP(&conf.Editor, "editor", "e", editor, "Selection of editor")
	rootCmd.Flags().StringVarP(&conf.DirPath, "notes-directory", "d", userHome, "Sets directory of main notes")
}

func rootRun(cmd *cobra.Command, args []string) {
	if err := gonoted.New(&conf); err != nil {
		cmd.PrintErrln(err)
	}
}

const (
	usage = `
	gonote new <note-name>
	gonote ls
	`

	defaultEditor = "nano"
)
