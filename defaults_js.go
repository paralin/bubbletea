//go:build js
// +build js

package tea

import (
	"io"
	"os"
)

// DisableConsole globally disables the console.
var DisableConsole bool = true

// DefaultInput is the default input used.
var DefaultInput io.Reader

// DefaultInputWriter is a writer for DefaultInput on js.
var DefaultInputWriter io.Writer

func init() {
	DefaultInput, DefaultInputWriter, _ = os.Pipe()
}

// DefaultOutput is the default output used.
var DefaultOutput io.Writer = io.Discard
