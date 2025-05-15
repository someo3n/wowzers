// layout is a package containing everything about layouts in wowzers
package layout

// Arranger is a type that defines function that needs to take in
// a layout rendering Context and output an Element. This is used
// when arranging is needed
type Arranger func(ctx *Context) Element

// Context is the context of the layout rendering
type Context struct {
	// Masters is the slice of master clients
	Masters []*Client

	// Slaves is the slice of slave clients
	Slaves []*Client

	// Clients is the slice of all the clients
	Clients []*Client

	// Focused is the client index that is focused
	Focused uint
}

// Layout is a struct that stores layout metadata and an arranging function
type Layout struct {
	// Name is the name of the layout
	Name string

	// RearrangeOnFocusChange makes the layout get rearranged when focus changes
	RearrangeOnFocusChange bool

	// Arrange is the callback function that gets invoked when a rearrange is needed
	Arrange Arranger
}
