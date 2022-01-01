//go:build noconsole
// +build noconsole

package tea

func init() {
	DisableConsole = true
}
