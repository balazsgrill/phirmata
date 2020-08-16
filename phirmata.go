package phirmata

import (
	"io"

	"github.com/argandas/goduino/firmata"
	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/i2c"
)

// Dev provides access to functionalities
type Dev interface {
	Pins() []gpio.PinIO
	I2C() i2c.Bus
}

type dev struct {
	pins []gpio.PinIO
	i2c  i2c.Bus
}

// New attempts to connect to a device over the given serial connection using the Firmata protocol
func New(conn io.ReadWriteCloser) (Dev, error) {
	f := firmata.New()
	err := f.Connect(conn)
	if err != nil {
		return nil, err
	}
	pincount := len(f.Pins())
	pins := make([]gpio.PinIO, pincount)
	for i := 0; i < pincount; i++ {
		p, err := NewPin(f, i)
		if err != nil {
			return nil, err
		}
		pins[i] = p
	}
	return &dev{
		pins: pins,
		i2c:  nil,
	}, nil
}

func (d *dev) Pins() []gpio.PinIO {
	return d.pins
}
func (d *dev) I2C() i2c.Bus {
	return d.i2c
}
