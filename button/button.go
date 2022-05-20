package button

import (
	"machine"
)

func New() machine.Pin {
	button := machine.D1
	button.Configure(machine.PinConfig{Mode: machine.PinInput})
	return button
}
