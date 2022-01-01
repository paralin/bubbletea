//go:build !js
// +build !js

package tea

import (
	"io"
	"os"
)

// DisableConsole globally disables the console.
var DisableConsole bool

// DefaultInput is the default input used.
var DefaultInput io.Reader = os.Stdin

// DefaultOutput is the default output used.
var DefaultOutput io.Writer = os.Stdout
