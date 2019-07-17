//--------------------------------------------------------------------------------------------------
//
// Copyright (c) 2018 Denis Dyakov
//    poritons Copyright (c) 2019 Iron Heart Consulting, LLC
// Adapted for ADS1115 device by Steve Conklin
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

// Package implements controlling the A/D and reading sampled values for the ADS1115 A/D Converter
package ads

import (
	"github.com/sconklin/go-i2c"
)

// SensorType identify which Bosch Sensortec
// temperature and pressure sensor is used.
// BMP180 and BMP280 are supported.
type SensorType int

// Implement Stringer interface.
func (v SensorType) String() string {
	if v == ADS1115 {
		return "ADS1115"
	}
	return "!!! unknown !!!"
}

const (
	// ADS1115 A/D Converter
	ADS1115 SensorType = iota
)

// InputMuxMode : Input Multiplexer Mode
const (
	MUX_DIFFERENTIAL_0_1 = 0 // Differential, 0 Positive, 1 Negative
	MUX_DIFFERENTIAL_0_3 = 1 // Differential, 0 Positive, 3 Negative
	MUX_DIFFERENTIAL_1_3 = 2 // Differential, 1 Positive, 3 Negative
	MUX_DIFFERENTIAL_2_3 = 3 // Differential, 2 Positive, 3 Negative
	MUX_SINGLE_0         = 4 // Single Ended 0
	MUX_SINGLE_1         = 5 // Single Ended 1
	MUX_SINGLE_2         = 6 // Single Ended 2
	MUX_SINGLE_3         = 7 // Single Ended 3
	MUX_MAX              = 7
)

// PGAMode : Programmable Gain Amplifier config
const (
	PGA_6_144  = 0 // Full Scale Range = +/- 6.144V
	PGA_4_096  = 1 // Full Scale Range = +/- 4.096V
	PGA_2_048  = 2 // Full Scale Range = +/- 2.048V
	PGA_1_024  = 3 // Full Scale Range = +/- 1.024V
	PGA_0_512  = 4 // Full Scale Range = +/- 0.512V
	PGA_0_256  = 5 // Full Scale Range = +/- 0.128V
	PGA_0_256a = 6 // Full Scale Range = +/- 0.128V
	PGA_0_256b = 7 // Full Scale Range = +/- 0.128V
	PGA_MAX    = 7
)

// Mode : Conversion Mode
const (
	MODE_CONTINUOUS  = 0 // Continuous Conversion
	MODE_SINGLE_SHOT = 1 // Single Shot Conversion
	MODE_MAX         = 1
)

// Datarate is the A/D sampling rate
const (
	RATE_8   = 0 // 8 Samples per Second
	RATE_16  = 1 // 16 Samples per Second
	RATE_32  = 2 // 32 Samples per Second
	RATE_64  = 3 // 64 Samples per Second
	RATE_128 = 4 // 128 Samples per Second
	RATE_150 = 5 // 150 Samples per Second
	RATE_475 = 6 // 475 Samples per Second
	RATE_860 = 7 // 860 Samples per Second
	RATE_MAX = 7
)

// Comparator Mode
const (
	COMP_MODE_TRADITIONAL = 0
	COMP_MODE_WINDOW      = 1
	COMP_MODE_MAX         = 1
)

// Comparator Polarity
const (
	COMP_POL_ACTIVE_LOW  = 0
	COMP_POL_ACTIVE_HIGH = 1
	COMP_POL_MAX         = 1
)

// Comparator Latch
const (
	COMP_LAT_OFF = 0
	COMP_LAT_ON  = 1
	COMP_LAT_MAX = 1
)

// Comparator Queue
const (
	COMP_QUE_ONE     = 0
	COMP_QUE_TWO     = 1
	COMP_QUE_FOUR    = 2
	COMP_QUE_DISABLE = 3
	COMP_QUE_MAX     = 3
)

// SensorInterface is an Abstract ADSx sensor interface
type SensorInterface interface {

	// ReadConfig reads configuration from the chip
	ReadConfig(i2c *i2c.I2C) (uint16, error)

	// WriteConfig writes the stored configuration to the chip
	WriteConfig(i2c *i2c.I2C) error

	// SetMuxMode sets the stored configuration (does not write to chip)
	SetMuxMode(uint16) error

	// SetPGAMode sets the stored configuration (does not write to chip)
	SetPgaMode(uint16) error

	// SetConversionMode sets the stored configuration (does not write to chip)
	SetConversionMode(uint16) error

	// SetDataRate sets the stored configuration (does not write to chip)
	SetDataRate(uint16) error

	// SetComparatorMode sets the stored configuration (does not write to chip)
	SetComparatorMode(uint16) error

	// SetComparatorPolarity sets the stored configuration (does not write to chip)
	SetComparatorPolarity(uint16) error

	// SetComparatorLatch sets the stored configuration (does not write to chip)
	SetComparatorLatch(uint16) error

	// SetComparatorQueue sets the stored configuration (does not write to chip)
	SetComparatorQueue(uint16) error

	// ReadStatus reads the status register from the chip
	// returns nonzero if a conversion is in progress
	ReadStatus(i2c *i2c.I2C) (uint16, error)

	// StartConversion will start a conversion in single-shot mode
	StartConversion(i2c *i2c.I2C) error

	// ReadLoThreshold reads the low comparator threshold from the chip
	ReadLoThreshold(i2c *i2c.I2C) (int16, error)

	// ReadHiThreshold reads the high comparator threshold from the chip
	ReadHiThreshold(i2c *i2c.I2C) (int16, error)

	// ReadConversion reads the converted value from the chip
	ReadConversion(i2c *i2c.I2C) (int16, error)
}

// ADS represents only one model of A/D so far
type ADS struct {
	sensorType SensorType
	i2c        *i2c.I2C
	ads        SensorInterface
}

// NewADS creats a new device interface
func NewADS(sensorType SensorType, i2c *i2c.I2C) (*ADS, error) {
	v := &ADS{sensorType: sensorType, i2c: i2c}
	switch sensorType {
	case ADS1115:
		v.ads = &SensorADS1115{}
	}

	_, err := v.ads.ReadConfig(i2c)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// ReadConfig from the chip
func (v *ADS) ReadConfig() (uint16, error) {
	t, err := v.ads.ReadConfig(v.i2c)
	return t, err
}

// WriteConfig to the chip from the stored config data
func (v *ADS) WriteConfig() error {
	err := v.ads.WriteConfig(v.i2c)
	return err
}

// SetMuxMode in stored config
func (v *ADS) SetMuxMode(imm uint16) error {
	err := v.ads.SetMuxMode(imm)
	return err
}

// SetPgaMode in stored config
func (v *ADS) SetPgaMode(pm uint16) error {
	err := v.ads.SetPgaMode(pm)
	return err
}

// SetConversionMode in stored config
func (v *ADS) SetConversionMode(md uint16) error {
	err := v.ads.SetConversionMode(md)
	return err
}

// SetDataRate in stored config
func (v *ADS) SetDataRate(dr uint16) error {
	err := v.ads.SetDataRate(dr)
	return err
}

// SetComparatorMode in stored config
func (v *ADS) SetComparatorMode(cm uint16) error {
	err := v.ads.SetComparatorMode(cm)
	return err
}

//	SetComparatorPolarity in stored config
func (v *ADS) SetComparatorPolarity(cp uint16) error {
	err := v.ads.SetComparatorPolarity(cp)
	return err
}

//	SetComparatorLatch in stored config
func (v *ADS) SetComparatorLatch(cl uint16) error {
	err := v.ads.SetComparatorLatch(cl)
	return err
}

//	SetComparatorQueue in stored config
func (v *ADS) SetComparatorQueue(cq uint16) error {
	err := v.ads.SetComparatorQueue(cq)
	return err
}

// ReadStatus from the chip
func (v *ADS) ReadStatus() (uint16, error) {
	t, err := v.ads.ReadStatus(v.i2c)
	return t, err
}

// StartConversion if in single-shot mode
func (v *ADS) StartConversion() error {
	err := v.ads.StartConversion(v.i2c)
	return err
}

// ReadLoThreshold for comparator from the chip
func (v *ADS) ReadLoThreshold() (int16, error) {
	t, err := v.ads.ReadLoThreshold(v.i2c)
	return t, err
}

// ReadHiThreshold for comparator from the chip
func (v *ADS) ReadHiThreshold() (int16, error) {
	t, err := v.ads.ReadHiThreshold(v.i2c)
	return t, err
}

// ReadConversion value from the chip
func (v *ADS) ReadConversion() (int16, error) {
	t, err := v.ads.ReadConversion(v.i2c)
	return t, err
}
