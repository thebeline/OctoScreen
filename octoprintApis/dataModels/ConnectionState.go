package dataModels

import (
	"strings"
)


type ConnectionState string

const (
	Operational ConnectionState = "Operational"
)

// The states are based on:
// https://github.com/foosel/OctoPrint/blob/77753ca02602d3a798d6b0a22535e6fd69ff448a/src/octoprint/util/comm.py#L549

func (s ConnectionState) IsOperational() bool {
	return strings.HasPrefix(string(s), "Operational")
}

func (s ConnectionState) IsPrinting() bool {
	return strings.HasPrefix(string(s), "Printing") ||
		strings.HasPrefix(string(s), "Starting") ||
		strings.HasPrefix(string(s), "Sending") ||
		strings.HasPrefix(string(s), "Paused") ||
		strings.HasPrefix(string(s), "Pausing") ||
		strings.HasPrefix(string(s), "Transfering")
}

func (s ConnectionState) IsOffline() bool {
	return strings.HasPrefix(string(s), "Offline") ||
		strings.HasPrefix(string(s), "Closed")
}

func (s ConnectionState) IsError() bool {
	return strings.HasPrefix(string(s), "Error") ||
		strings.HasPrefix(string(s), "Unknown")
}

func (s ConnectionState) IsConnecting() bool {
	return strings.HasPrefix(string(s), "Opening") ||
		strings.HasPrefix(string(s), "Detecting") ||
		strings.HasPrefix(string(s), "Connecting") ||
		strings.HasPrefix(string(s), "Detecting")
}
