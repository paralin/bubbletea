module github.com/charmbracelet/bubbletea

go 1.16

replace github.com/containerd/console => github.com/paralin/containerd-console v1.0.4-0.20220116023838-20ef04a44192 // shim-js

require (
	github.com/containerd/console v1.0.3
	github.com/mattn/go-isatty v0.0.14
	github.com/muesli/ansi v0.0.0-20211031195517-c9f0611b6c70
	github.com/muesli/reflow v0.3.0
	github.com/muesli/termenv v0.10.0
	golang.org/x/sys v0.0.0-20220114195835-da31bd327af9
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211
)
