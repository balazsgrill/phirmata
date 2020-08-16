package phirmata

import (
	"fmt"
	"time"

	"github.com/argandas/goduino/firmata"
	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/physic"
)

type pin struct {
	owner *firmata.Firmata
	pin   int
}

// NewPin creates a Periph GPIO handle for the given Firmata Pin
func NewPin(owner *firmata.Firmata, pinid int) (gpio.PinIO, error) {
	return &pin{
		owner: owner,
		pin:   pinid,
	}, nil
}

func (p *pin) In(pull gpio.Pull, edge gpio.Edge) error {
	if err := p.owner.SetPinMode(p.pin, firmata.Input); err != nil {
		return err
	}
	if err := p.owner.ReportDigital(p.pin, 1); err != nil {
		return err
	}
	return nil
}

func (p *pin) Read() gpio.Level {
	if 0 == p.owner.Pins()[p.pin].Value {
		return gpio.Low
	}
	return gpio.High
}

func (*pin) WaitForEdge(timeout time.Duration) bool {
	// TODO support events
	return false
}

func (*pin) Pull() gpio.Pull {
	return gpio.Float
}

func (*pin) DefaultPull() gpio.Pull {
	return gpio.Float
}

func (p *pin) Out(l gpio.Level) error {
	fpin := p.owner.Pins()[p.pin]
	if fpin.Mode != firmata.Output {
		if err := p.owner.SetPinMode(p.pin, firmata.Output); err != nil {
			return err
		}
	}
	if err := p.owner.ReportDigital(p.pin, 0); err != nil {
		return err
	}
	v := 0
	if l == gpio.High {
		v = 1
	}
	return p.owner.DigitalWrite(p.pin, v)
}
func (p *pin) PWM(duty gpio.Duty, f physic.Frequency) error {
	// Duty cycle is scaled from 24 bits to 8 bits
	return p.owner.AnalogWrite(p.pin, int(duty>>16))
}

func (p *pin) Name() string {
	return fmt.Sprintf("Firmata_%d", p.pin)
}

func (p *pin) Number() int {
	return p.pin
}

func (*pin) Function() string {
	// TODO unimplemenred
	return string(gpio.IN)
}
func (p *pin) String() string {
	return p.Name()
}
func (*pin) Halt() error {
	// TODO unimplemented
	return nil
}
