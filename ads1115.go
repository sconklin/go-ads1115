//--------------------------------------------------------------------------------------------------
//
// Copyright (c) 2018 Denis Dyakov
// Adapted for ads1115 by Steve Conklin
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and
// associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all copies or substantial
// portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING
// BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
//
//--------------------------------------------------------------------------------------------------

package ads

import (
	"errors"
	i2c "github.com/sconklin/go-i2c"
)

// ads1115 A/D memory map
const (
	ADS1115_ADDR_CONVERSION = 0
	ADS1115_ADDR_CONFIG     = 1
	ADS1115_ADDR_LO_THRESH  = 2
	ADS1115_ADDR_HI_THRESH  = 3
)

// SensorADS1115 contains internal register config items
type SensorADS1115 struct {
	ads                      SensorInterface
	MuxConfig                uint16
	PGAConfig                uint16
	ModeConfig               uint16
	DataRateConfig           uint16
	ComparatorModeConfig     uint16
	ComparatorPolarityConfig uint16
	ComparatorLatchConfig    uint16
	ComparatorQueueConfig    uint16
}

// Static cast to verify at compile time
// that type implement interface.
var _ SensorInterface = &SensorADS1115{}

var config SensorADS1115

// ReadConfig reads the configuration register
func (v *SensorADS1115) ReadConfig(i2c *i2c.I2C) (uint16, error) {
	config, err := i2c.ReadRegU16BE(ADS1115_ADDR_CONFIG)
	if err != nil {
		return 0, err
	}
	return config, nil
}

// WriteConfig sets the config register from the stored config
func (v *SensorADS1115) WriteConfig(i2c *i2c.I2C) error {
	var cfgdata uint16

	cfgdata = config.ComparatorQueueConfig & 0x03
	cfgdata |= (config.ComparatorLatchConfig & 1) << 2
	cfgdata |= (config.ComparatorPolarityConfig & 1) << 3
	cfgdata |= (config.ComparatorModeConfig & 1) << 4
	cfgdata |= (config.DataRateConfig & 0x07) << 5
	cfgdata |= (config.ModeConfig & 1) << 8
	cfgdata |= (config.PGAConfig & 0x07) << 9
	cfgdata |= (config.MuxConfig & 0x07) << 12

	err := i2c.WriteRegU16BE(ADS1115_ADDR_CONFIG, cfgdata)
	if err != nil {
		return err
	}
	return nil
}

// SetMuxMode sets the stored config
func (v *SensorADS1115) SetMuxMode(imm uint16) error {
	if imm > MUX_MAX {
		return errors.New("Invalid value for Mux Mode")
	}
	config.MuxConfig = imm
	return nil
}

// SetPgaMode in config
func (v *SensorADS1115) SetPgaMode(pm uint16) error {
	if pm > PGA_MAX {
		return errors.New("Invalid value for PGA Mode")
	}
	config.PGAConfig = pm
	return nil
}

// SetConversionMode in stored config
func (v *SensorADS1115) SetConversionMode(md uint16) error {
	if md > MODE_MAX {
		return errors.New("Invalid value for Conversion Mode")
	}
	config.ModeConfig = md
	return nil
}

// SetDataRate in stored config
func (v *SensorADS1115) SetDataRate(dr uint16) error {
	if dr > RATE_MAX {
		return errors.New("Invalid value for Data Rate")
	}
	config.DataRateConfig = dr
	return nil
}

// SetComparatorMode in stored config
func (v *SensorADS1115) SetComparatorMode(cm uint16) error {
	if cm > COMP_MODE_MAX {
		return errors.New("Invalid value for Comparator Mode")
	}
	config.ComparatorModeConfig = cm
	return nil
}

//	SetComparatorPolarity in stored config
func (v *SensorADS1115) SetComparatorPolarity(cp uint16) error {
	if cp > COMP_POL_MAX {
		return errors.New("Invalid value for Comparator Polarity")
	}
	config.ComparatorPolarityConfig = cp
	return nil
}

//	SetComparatorLatch in stored config
func (v *SensorADS1115) SetComparatorLatch(cl uint16) error {
	if cl > COMP_LAT_MAX {
		return errors.New("Invalid value for Comparator Latch")
	}
	config.ComparatorLatchConfig = cl
	return nil
}

//	SetComparatorQueue in stored config
func (v *SensorADS1115) SetComparatorQueue(cq uint16) error {
	if cq > COMP_QUE_MAX {
		return errors.New("Invalid value for Comparator Queue")
	}
	config.ComparatorQueueConfig = cq
	return nil
}

// ReadStatus from the chip, returns nonzero if a conversion is in progress
func (v *SensorADS1115) ReadStatus(i2c *i2c.I2C) (uint16, error) {
	t, err := v.ads.ReadConfig(i2c)
	if err != nil {
		return 0, err
	}
	t = t | 0x80
	return t, err
}

// StartConversion if in single-shot mode
func (v *SensorADS1115) StartConversion(i2c *i2c.I2C) error {
	cfg, err := i2c.ReadRegU16BE(ADS1115_ADDR_CONFIG)
	if err != nil {
		return err
	}
	cfg = cfg | 0x80
	err = i2c.WriteRegU16BE(ADS1115_ADDR_CONFIG, cfg)
	if err != nil {
		return err
	}
	return nil
}

// ReadLoThreshold from the chip
func (v *SensorADS1115) ReadLoThreshold(i2c *i2c.I2C) (int16, error) {
	t, err := i2c.ReadRegS16BE(ADS1115_ADDR_LO_THRESH)
	if err != nil {
		return 0, err
	}
	return t, err
}

// ReadHiThreshold from the chip
func (v *SensorADS1115) ReadHiThreshold(i2c *i2c.I2C) (int16, error) {
	t, err := i2c.ReadRegS16BE(ADS1115_ADDR_HI_THRESH)
	if err != nil {
		return 0, err
	}
	return t, err
}

// ReadConversion value from the chip
func (v *SensorADS1115) ReadConversion(i2c *i2c.I2C) (int16, error) {
	t, err := i2c.ReadRegS16BE(ADS1115_ADDR_CONVERSION)
	if err != nil {
		return 0, err
	}
	return t, err
}
