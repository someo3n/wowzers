// client is a package containing everything about client management in wowzers
package client

// Manager is a manager of Clients
type Manager struct {
	// TagMask is the current tag mask
	TagMask TagMask

	// Focused is the handle of the currently focused client
	Focused Handle

	// Clients is a slice of all managed clients
	Clients []*Client
}
