NOTE: This is an in-progress port of the bsbmp library to the ADS1115

It's a mess until this warning is removed

=======================================================================================================

Texas Instruments ADS1115 four channel A/d
=============================================================================================

[![Build Status](https://travis-ci.org/sconklin/go-bsbmp.svg?branch=master)](https://travis-ci.org/sconklin/go-ads1115)
[![Go Report Card](https://goreportcard.com/badge/github.com/sconklin/go-ads1115)](https://goreportcard.com/report/github.com/sconklin/go-ads1115)
[![GoDoc](https://godoc.org/github.com/sconklin/go-ads1115?status.svg)](https://godoc.org/github.com/sconklin/go-ads1115)
[![MIT License](http://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)

ADS1115 ([pdf reference](https://raw.github.com/sconklin/go-ads1115/master/docs/ads1115.pdf)), is an A/D converter used by Arduino and Raspberry PI developers.
The A/D works with the i2c bus interface:
![image](https://raw.github.com/sconklin/go-ads1115/master/docs/adafruit_1085.jpg)

This is a library written in [Go programming language](https://golang.org/) for Raspberry PI and counterparts, which gives you the ability to program the A/D and read the data. (Handling all necessary i2c-bus communication).

Golang usage
------------


```go
func main() {
	// Create new connection to i2c-bus on 1 line with address 0x76.
	// Use i2cdetect utility to find device address over the i2c-bus
	i2c, err := i2c.NewI2C(0x48, 1)
	if err != nil {
		log.Fatal(err)
	}
	defer i2c.Close()
	// Uncomment next line to supress verbose output
	//logger.ChangePackageLogLevel("i2c", logger.InfoLevel)

	sensor, err := ads.NewADS(ads.ADS1115, i2c) // signature=0x58

	if err != nil {
		lg.Fatal(err)
	}

	config, err := sensor.ReadConfig()
	if err != nil {
		lg.Fatal(err)
	}
	lg.Infof("This A/D has initial config: 0x%x", config)

	err = sensor.SetMuxMode(ads.MUX_SINGLE_0)
	if err != nil {
		lg.Fatal(err)
	}
	lg.Infof("  Configured for Single Ended Channel 0")

	err = sensor.SetPgaMode(ads.PGA_0_256)
	if err != nil {
		lg.Fatal(err)
	}
	lg.Infof("  Configured for +/- 128 mV Full Scale")

	err = sensor.SetConversionMode(ads.MODE_CONTINUOUS)
	if err != nil {
		lg.Fatal(err)
	}
	lg.Infof("  Configured for continuous sampling")

	err = sensor.SetDataRate(ads.RATE_8)
	if err != nil {
		lg.Fatal(err)
	}
	lg.Infof("  Configured for 8 Samples per Second")

	err = sensor.SetComparatorMode(ads.COMP_MODE_TRADITIONAL)
	if err != nil {
		lg.Fatal(err)
	}
	lg.Infof("  Configured for traditional comparator mode")

	err = sensor.SetComparatorPolarity(ads.COMP_POL_ACTIVE_LOW)
	if err != nil {
		lg.Fatal(err)
	}
	lg.Infof("  Configured comparator active low")

	err = sensor.SetComparatorLatch(ads.COMP_LAT_OFF)
	if err != nil {
		lg.Fatal(err)
	}
	lg.Infof("  Configured comparator latch off")

	err = sensor.SetComparatorQueue(ads.COMP_QUE_DISABLE)
	if err != nil {
		lg.Fatal(err)
	}
	lg.Infof("  Configured comparator queue disabled")

	err = sensor.WriteConfig()
	if err != nil {
		lg.Fatal(err)
	}
	lg.Infof("  Wrote new Config to A/D")

	config, err = sensor.ReadConfig()
	if err != nil {
		lg.Fatal(err)
	}
	lg.Infof("This A/D has final config: 0x%x", config)

	for i := 1; i < 5; i++ {
		time.Sleep(2 * time.Second)
		val, err := sensor.ReadConversion()
		if err != nil {
			lg.Fatal(err)
		}
		lg.Infof("A/D value: 0x%x", val)
	}
}
    
```

Getting help
------------

GoDoc [documentation](http://godoc.org/github.com/sconklin/go-ads1115)

Installation
------------

```bash
$ go get -u github.com/sconklin/go-ads1115
```

Troubleshooting
--------------

- *How to obtain fresh Golang installation to RPi device (either any RPi clone):*
If your RaspberryPI golang installation taken by default from repository is outdated, you may consider
to install actual golang mannualy from official Golang [site](https://golang.org/dl/). Download
tar.gz file containing armv6l in the name. Follow installation instructions.

- *How to enable I2C bus on RPi device:*
If you employ RaspberryPI, use raspi-config utility to activate i2c-bus on the OS level.
Go to "Interfaceing Options" menu, to active I2C bus.
Probably you will need to reboot to load i2c kernel module.
Finally you should have device like /dev/i2c-1 present in the system.

- *How to find I2C bus allocation and device address:*
Use i2cdetect utility in format "i2cdetect -y X", where X may vary from 0 to 5 or more,
to discover address occupied by peripheral device. To install utility you should run
`apt install i2c-tools` on debian-kind system. `i2cdetect -y 1` sample output:
	```
	     0  1  2  3  4  5  6  7  8  9  a  b  c  d  e  f
	00:          -- -- -- -- -- -- -- -- -- -- -- -- --
	10: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
	20: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
	30: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
	40: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
	50: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
	60: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
	70: -- -- -- -- -- -- 76 --    
	```

Contributing authors
------------------

* [Denis Dyakov](https://github.com/d2r2): Original sensor go library for the BMP388, with associated I2C and logging libraries
* [Steve Conklin](https://github.com/sconklin): Port of the library for ads-1115.


Contact
-------

Please use [Github issue tracker](https://github.com/sconklin/go-ads1115/issues) for filing bugs or feature requests.


License
-------

Go-ads1115 is licensed under MIT License.
