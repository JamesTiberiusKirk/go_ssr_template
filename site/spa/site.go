package spa

type Site struct {
	MenuID string // TODO: Possible future integration with the frame meny in the future
	// Title string
	// Frame bool // TODO: figure out how to implement the frame...think will need to make index from a template instead
	Path  string
	Dist  string
	Index string
}
