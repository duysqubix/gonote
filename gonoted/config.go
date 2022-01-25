package gonoted

import (
	"log"

	"github.com/chigopher/pathlib"
)

// Config holds gonote settings
type Config struct {
	// selection of editor - defaults to $EDITOR
	Editor string

	// path to note directory
	DirPath string

	// default file type
	FileType string
}

func (c *Config) DirPathObj() *pathlib.Path {
	d := MakePathObj(c.DirPath)

	if exists, _ := d.Exists(); !exists {
		if err := d.MkdirAll(); err != nil {
			log.Fatal(err)
		}
	}

	return d
}
