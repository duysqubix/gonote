package gonoted

// Config holds gonote settings
type Config struct {
	// selection of editor - defaults to $EDITOR
	Editor string

	// path to note directory
	DirPath string
}
