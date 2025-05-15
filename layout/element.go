// layout is a package containing everything about layouts in wowzers
package layout

import "wowzers/client"

// ElementType is the type of the layout element idk
type ElementType uint8

const (
	// ClientType is the type that Client elements use
	ClientType ElementType = iota

	// OverlayType is the type for Collection elements that means to
	// put each child on top of eachother in the Collections area
	OverlayType

	// VSplitType is the type for Collection elements that means to
	// split the Collections area vertically and then put each child there
	VSplitType

	// HSplitType is the type for Collection elements that means to
	// split the Collections area horizontally and then put each child there
	HSplitType
)

// Element is the base interface that is needed to make a layout
type Element interface {
	// Type returns the type of the Element
	Type() ElementType
}

// Client here is an layout Element that means to put a Client in
// the elements base
type Client struct {
	// Handle is the handle of the client
	Handle client.Handle
}

// Type returns the type of the Client Element which is ClientType
func (c *Client) Type() ElementType {
	return ClientType
}

// NewClient creates a new Client element with that client handle
func NewClient(handle client.Handle) *Client {
	return &Client{Handle: handle}
}

// Collection is an Element that has children
type Collection struct {
	// CollectionType is the type of the Collection.
	// Overlay, VSplit and HSplit are supported
	CollectionType ElementType

	// Children are the Children of the Element
	Children []Element
}

// Type returns the type of the Collection Element which is the
// value of its CollectionType
func (c *Collection) Type() ElementType {
	return c.CollectionType
}

// NewCollection creates a new Collection of that collection type and those children
func NewCollection(collectionType ElementType, children []Element) *Collection {
	return &Collection{
		CollectionType: collectionType,
		Children:       children,
	}
}

// NewOverlay creates a new Collection element with OverlayType type and those children
func NewOverlay(children []Element) *Collection {
	return NewCollection(OverlayType, children)
}

// NewOverlay creates a new Collection element with VSplitType type and those children
func NewVSplit(children []Element) *Collection {
	return NewCollection(VSplitType, children)
}

// NewOverlay creates a new Collection element with HSplitType type and those children
func NewHSplit(children []Element) *Collection {
	return NewCollection(HSplitType, children)
}
