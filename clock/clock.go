package clock

import (
	"machine"
	"strconv"
)

type Clock struct {
	time               string //25:00 = 25 minutes and 00 seconds
	hardwareSecondInMs int
	pinA               machine.Pin
	pinB               machine.Pin
	pinC               machine.Pin
	pinD               machine.Pin
	pinE               machine.Pin
	pinF               machine.Pin
	pinG               machine.Pin
	D1                 machine.Pin
	D2                 machine.Pin
	D3                 machine.Pin
	D4                 machine.Pin
}

func New() *Clock {
	clock := &Clock{}
	clock.time = "00:00"
	clock.hardwareSecondInMs = 1000;

	clock.pinA = machine.ADC1
	clock.pinA.Configure(machine.PinConfig{Mode: machine.PinOutput})

	clock.pinB = machine.ADC5
	clock.pinB.Configure(machine.PinConfig{Mode: machine.PinOutput})

	clock.pinC = machine.D10
	clock.pinC.Configure(machine.PinConfig{Mode: machine.PinOutput})

	clock.pinD = machine.D11
	clock.pinD.Configure(machine.PinConfig{Mode: machine.PinOutput})

	clock.pinE = machine.D12
	clock.pinE.Configure(machine.PinConfig{Mode: machine.PinOutput})

	clock.pinF = machine.ADC2
	clock.pinF.Configure(machine.PinConfig{Mode: machine.PinOutput})

	clock.pinG = machine.D13
	clock.pinG.Configure(machine.PinConfig{Mode: machine.PinOutput})

	clock.D1 = machine.ADC0
	clock.D1.Configure(machine.PinConfig{Mode: machine.PinOutput})

	clock.D2 = machine.ADC3
	clock.D2.Configure(machine.PinConfig{Mode: machine.PinOutput})

	clock.D3 = machine.ADC4
	clock.D3.Configure(machine.PinConfig{Mode: machine.PinOutput})

	clock.D4 = machine.D7
	clock.D4.Configure(machine.PinConfig{Mode: machine.PinOutput})

	return clock
}

func (c *Clock) SetHardwareSecondInMs(value int) {
  c.hardwareSecondInMs = value;
}

func (c *Clock) SetTime(time string) {
	c.time = time
}

func (c *Clock) GetTimeInt() (int, int) {
	minutes, _ := strconv.Atoi(c.time[0:2])
	seconds, _ := strconv.Atoi(c.time[3:5])
	return minutes, seconds
}

func (c *Clock) GetHardwareSecondInMs() int {
	return c.hardwareSecondInMs
}

func (c *Clock) DisplayDigit(digit int, digitPosition int) {
	switch digitPosition {
	case 0:
		c.D4.High()
	case 1:
		c.D3.High()
	case 2:
		c.D2.High()
	case 3:
		c.D1.High()
	}

	if digit != 1 && digit != 4 {
		c.pinA.Low()
	}
	if digit != 5 && digit != 6 {
		c.pinB.Low()
	}
	if digit != 2 {
		c.pinC.Low()
	}
	if digit != 1 && digit != 4 && digit != 7 {
		c.pinD.Low()
	}
	if digit == 2 || digit == 6 || digit == 8 || digit == 0 {
		c.pinE.Low()
	}
	if digit != 1 && digit != 2 && digit != 3 && digit != 7 {
		c.pinF.Low()
	}
	if digit != 0 && digit != 1 && digit != 7 {
		c.pinG.Low()
	}
}

func (c *Clock) ResetPins() {
	c.D1.Low()
	c.D2.Low()
	c.D3.Low()
	c.D4.Low()

	c.pinA.High()
	c.pinB.High()
	c.pinC.High()
	c.pinD.High()
	c.pinE.High()
	c.pinF.High()
	c.pinG.High()
}
