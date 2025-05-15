// client is a package containing everything about client management in wowzers
package client

import "wowzers/util"

// TagMask is a bitmask for tags
type TagMask uint32

// Handle is a client handle
type Handle uint32

// Client is a managed client by our WM
type Client struct {
	// Handle is the client handle
	Handle

	// TagMask is the tag mask
	TagMask

	// Factor is how much space the client should take up
	// of its parent when in a layout
	Factor float32

	// IsFloating is a flag to tell if the client is floating
	IsFloating bool

	// The position and size for the client if its floating
	FloatingRect util.Rect
}
