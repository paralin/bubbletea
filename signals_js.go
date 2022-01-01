//go:build js
// +build js

package tea

import (
	"context"
	"os"
)

// listenForResize is not available on this platform
func listenForResize(ctx context.Context, output *os.File, msgs chan Msg,
	errs chan error, done chan struct{}) {
	close(done)
}
