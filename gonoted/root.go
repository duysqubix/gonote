package gonoted

import (
	"os/exec"
	"time"

	"github.com/chigopher/pathlib"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

// get name of quick note
func GetQuickNoteName() string {
	return "quicknote-" + time.Now().Format("02-01-2006")
}

// this assume fpath is valid path
func OpenEditorSync(cmd *cobra.Command, config *Config, fpath string) error {
	c := exec.Command(config.Editor, fpath)
	c.Stderr = cmd.ErrOrStderr()
	c.Stdin = cmd.InOrStdin()
	c.Stdout = cmd.OutOrStdout()

	if err := c.Run(); err != nil {
		return err
	}

	return nil
}

func MakePathObj(path string) *pathlib.Path {
	return pathlib.NewPathAfero(path, afero.NewOsFs())
}
