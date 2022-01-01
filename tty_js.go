//go:build js
// +build js

package tea

import (
	"errors"
	"os"
)

const ttySupported = false

func openInputTTY() (*os.File, error) {
	return nil, errors.New("tty not supported")
}

func (p *Program) initInput() error {
	return nil
}

func (p *Program) restoreInput() error {
	return nil
}
