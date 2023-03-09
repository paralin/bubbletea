module examples

go 1.13

replace github.com/containerd/console => github.com/paralin/containerd-console v1.0.4-0.20220116023838-20ef04a44192 // shim-js

replace github.com/charmbracelet/bubbletea => ../

require (
	github.com/charmbracelet/bubbles v0.13.0
	github.com/charmbracelet/bubbletea v0.22.0
	github.com/charmbracelet/glamour v0.5.0
	github.com/charmbracelet/lipgloss v0.5.0
	github.com/fogleman/ease v0.0.0-20170301025033-8da417bf1776
	github.com/lucasb-eyer/go-colorful v1.2.0
	github.com/mattn/go-isatty v0.0.17
	github.com/muesli/reflow v0.3.0
	github.com/muesli/termenv v0.15.0
)
