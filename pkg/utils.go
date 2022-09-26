package pkg

import (
	"context"
	"encoding/json"
	"github.com/joshuarubin/go-sway"
	"log"
	"os"
)

func ToStdOut(s any) {
	if err := json.NewEncoder(os.Stdout).Encode(s); err != nil {
		log.Fatal(err)
	}
}

func getCurrentLayout(ctx context.Context, c sway.Client) (sway.Input, bool) {
	inputs, err := c.GetInputs(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, input := range inputs {
		if input.XKBActiveLayoutName != nil {
			return input, true
		}
	}

	return sway.Input{}, false

}

func SetInitialLayout(ctx context.Context, c sway.Client, ev *EvHandler) {
	l, ok := getCurrentLayout(ctx, c)
	if !ok {
		return
	}

	ev.Input(ctx, sway.InputEvent{Input: l})

}
