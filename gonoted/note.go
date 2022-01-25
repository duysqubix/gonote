package gonoted

func New(c *Config) *Gonote {
	g := &Gonote{}

	return g
}

//
type Gonote struct {
}

func (g *Gonote) Run() error {
	return nil
}
