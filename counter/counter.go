package counter

import (
	"machine"
)

type Counter struct {
	value int
	pinA  machine.Pin
	pinB  machine.Pin
	pinC  machine.Pin
	pinD  machine.Pin
	pinE  machine.Pin
	pinF  machine.Pin
	pinG  machine.Pin
}

func New() *Counter {
	counter := &Counter{}
	counter.pinA = machine.D2
	counter.pinA.Configure(machine.PinConfig{Mode: machine.PinOutput})

	counter.pinB = machine.D3
	counter.pinB.Configure(machine.PinConfig{Mode: machine.PinOutput})

	counter.pinC = machine.D4
	counter.pinC.Configure(machine.PinConfig{Mode: machine.PinOutput})

	counter.pinD = machine.D5
	counter.pinD.Configure(machine.PinConfig{Mode: machine.PinOutput})

	counter.pinE = machine.D6
	counter.pinE.Configure(machine.PinConfig{Mode: machine.PinOutput})

	counter.pinF = machine.D8
	counter.pinF.Configure(machine.PinConfig{Mode: machine.PinOutput})

	counter.pinG = machine.D9
	counter.pinG.Configure(machine.PinConfig{Mode: machine.PinOutput})

	counter.ResetValue()
	return counter
}

func (c *Counter) Increment() {
	c.value++
	c.ResetPins()
	c.DisplayDigit(c.value)
}

func (c *Counter) ResetValue() {
	c.value = 0
	c.DisplayDigit(c.value)
}

func (c *Counter) DisplayDigit(digit int) {
	if digit != 1 && digit != 4 {
		c.pinA.High()
	}
	if digit != 5 && digit != 6 {
		c.pinB.High()
	}
	if digit != 2 {
		c.pinC.High()
	}
	if digit != 1 && digit != 4 && digit != 7 {
		c.pinD.High()
	}
	if digit == 2 || digit == 6 || digit == 8 || digit == 0 {
		c.pinE.High()
	}
	if digit != 1 && digit != 2 && digit != 3 && digit != 7 {
		c.pinF.High()
	}
	if digit != 0 && digit != 1 && digit != 7 {
		c.pinG.High()
	}
}

func (c *Counter) ResetPins() {
	c.pinA.Low()
	c.pinB.Low()
	c.pinC.Low()
	c.pinD.Low()
	c.pinE.Low()
	c.pinF.Low()
	c.pinG.Low()
}
