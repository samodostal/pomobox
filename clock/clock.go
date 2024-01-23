package clock

import (
	"machine"
	"strconv"
)

type Clock struct {
	time               string // 25:00 = 25 minutes and 00 seconds
	hardwareSecondInMs int
	spi                machine.SPI
}

func New() *Clock {
	clock := &Clock{}
	clock.time = "00:00"
	clock.hardwareSecondInMs = 1000

	clock.spi = machine.SPI{SCK: machine.D10, MOSI: machine.D11}

	clock.spi.Configure(machine.SPIConfig{
		Frequency: 1000000,
		Mode:      0,
	})

	return clock
}

func (c *Clock) SetHardwareSecondInMs(value int) {
	c.hardwareSecondInMs = value
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
	data := []byte{byte(digit), byte(digitPosition)}

	c.spi.Transfer(data, nil)
}
