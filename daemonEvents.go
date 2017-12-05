package common

// DaemonEvent codifies the manner in which the daemon can be moved into different statuses
type DaemonEvent int

// List of daemon events
const (
	DaemonEventRunCode DaemonEvent = iota
)
