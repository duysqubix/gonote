/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/chigopher/pathlib"
	"github.com/spf13/cobra"
)

type byPathObj []*pathlib.Path

func (b byPathObj) Len() int           { return len(b) }
func (b byPathObj) Less(i, j int) bool { return (*b[i]).Name() < (*b[j]).Name() }
func (b byPathObj) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Lists all notes by name and filters by file type extension",
	Run:   runLs,
}

func init() {
	rootCmd.AddCommand(lsCmd)
}

func runLs(cmd *cobra.Command, args []string) {

	dirPath := conf.DirPathObj()

	objs, _ := dirPath.Glob("*" + conf.FileType)

	sort.Sort(byPathObj(objs))

	fmt.Println("Notes:")
	for i, fobj := range objs {
		fname := fobj.Name()
		fname = strings.ReplaceAll(fname, "."+conf.FileType, "")
		fmt.Printf("%d) %s\n", i, fname)
	}
	fmt.Println()
}
