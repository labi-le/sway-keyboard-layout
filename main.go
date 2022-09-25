package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

type Input struct {
	Identifier           string   `json:"identifier"`
	Name                 string   `json:"name"`
	Vendor               int      `json:"vendor"`
	Product              int      `json:"product"`
	Type                 string   `json:"type"`
	XkbLayoutNames       []string `json:"xkb_layout_names,omitempty"`
	XkbActiveLayoutIndex int      `json:"xkb_active_layout_index,omitempty"`
	XkbActiveLayoutName  string   `json:"xkb_active_layout_name,omitempty"`

	LibInput LibInput `json:"libinput"`

	ScrollFactor float32 `json:"scroll_factor,omitempty"`
}

type LibInput struct {
	SendEvents      string  `json:"send_events"`
	Tap             string  `json:"tap,omitempty"`
	TapButtonMap    string  `json:"tap_button_map,omitempty"`
	TapDrag         string  `json:"tap_drag,omitempty"`
	TapDragLock     string  `json:"tap_drag_lock,omitempty"`
	AccelSpeed      float32 `json:"accel_speed,omitempty"`
	AccelProfile    string  `json:"accel_profile,omitempty"`
	NaturalScroll   string  `json:"natural_scroll,omitempty"`
	LeftHanded      string  `json:"left_handed,omitempty"`
	ClickMethod     string  `json:"click_method,omitempty"`
	MiddleEmulation string  `json:"middle_emulation,omitempty"`
	ScrollMethod    string  `json:"scroll_method,omitempty"`
	Dwt             string  `json:"dwt,omitempty"`
	ScrollButton    int     `json:"scroll_button,omitempty"`
}

type Output struct {
	Language string `json:"text"`
}

//goland:noinspection GoSnakeCaseUsage
const DEFAULT_UPDATE_INTERVAL = 1 * time.Second

func main() {
	Layout()

}

func getUpdateInterval() time.Duration {
	i := os.Getenv("SWAY_KEYBOARD_LAYOUT_UPDATE_INTERVAL")
	if i == "" {
		return DEFAULT_UPDATE_INTERVAL
	}

	t, err := strconv.Atoi(i)
	if err != nil {
		return DEFAULT_UPDATE_INTERVAL
	}

	return time.Duration(t) * time.Second
}

func Layout() {
	var (
		out bytes.Buffer
		err error
	)
	for {
		cmd := exec.Command("swaymsg", "-r", "-t", "get_inputs")
		cmd.Stdout = &out

		if err = cmd.Run(); err != nil {
			log.Fatal(err)
		}

		var inputs []Input
		if err = json.NewDecoder(&out).Decode(&inputs); err != nil {
			log.Fatal(err)
		}

		if layout, err := getFirstActiveKbdLayout(inputs); err != nil {
			log.Fatal(err)
		} else {
			if err := json.NewEncoder(os.Stdout).Encode(layout); err != nil {
				log.Fatal(err)
			}
		}

		time.Sleep(getUpdateInterval())
	}
}

func getFirstActiveKbdLayout(inputs []Input) (Output, error) {
	for _, input := range inputs {
		if input.XkbActiveLayoutName != "" {
			return Output{Language: input.XkbActiveLayoutName}, nil
		}
	}

	return Output{}, errors.New("no active keyboard layout found")
}
