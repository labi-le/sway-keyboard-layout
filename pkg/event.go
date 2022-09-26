package pkg

import (
	"context"
	"github.com/joshuarubin/go-sway"
)

type EvHandler struct {
	sway.EventHandler
}

func (e EvHandler) Input(_ context.Context, event sway.InputEvent) {
	if event.Input.XKBActiveLayoutName != nil {
		ToStdOut(Output{
			Text:    *event.Input.XKBActiveLayoutName,
			Tooltip: "Current keyboard layout",
			Class:   "keyboard-layout",
		})
	}
}
