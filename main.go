package main

import (
	"context"
	"github.com/joshuarubin/go-sway"
	"log"
	"sway-keyboard-layout/pkg"
)

// Turns out. ipc is more convenient than use exec.Command("swaymsg", "-t", "get_inputs")
func main() {
	var (
		ev  = &pkg.EvHandler{}
		ctx = context.Background()
	)

	client, err := sway.New(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// need to know the initial state of the keyboard layout
	// this is a trick to get the initial state
	pkg.SetInitialLayout(ctx, client, ev)

	if err := sway.Subscribe(ctx, ev, sway.EventTypeInput); err != nil {
		log.Fatal(err)
	}
}
