Bosch Sensortec BMP180, BMP280 temperature and pressure sensors
===============================================================

BMP180 and BMP280 are populare sensors among Arduino and Raspberry PI developers.
Both sensors are small and quite accurate working via i2c bus interface: (photos)

Here is a code written in [Go programming language](https://golang.org/) for Raspberry PI and counterparts, which gives you in the output temperature and atmosphere pressure values (making all necessary signal processing via i2c-bus behind the scene).

Golang usage
---------------


```
func main() {
	// Use i2cdetect utility to find device address at i2c-bus
	i2c, err := i2c.NewI2C(0x76, 1)
	if err != nil {
		log.Fatal(err)
	}
	defer i2c.Close()
	// Uncomment to seee verbose output
	// i2c.SetDebug(true)
	sensor, err := bsbmp.NewBMP(bsbmp.BMP280_TYPE, i2c)
	if err != nil {
		log.Fatal(err)
	}
	// Uncomment to seee verbose output
	// sensor.SetDebug(true)

	// Read temperature in celcius degree
	t, err := sensor.ReadTemperatureC(bsbmp.ACCURACY_STANDARD)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Temprature = %v*C\n", t)
	// Read atmosphere pressure in pascal
	p, err := sensor.ReadPressurePa(bsbmp.ACCURACY_STANDARD)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Pressure = %v Pa\n", p)
	// Read atmosphere pressure in mmHg
	p1, err := sensor.ReadPressureMmHg(bsbmp.ACCURACY_STANDARD)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Pressure = %v mmHg\n", p1)
	// Read atmosphere altitude in meters
	a, err := sensor.ReadAltitude(bsbmp.ACCURACY_STANDARD)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Altitude = %v m\n", a)
}
```

Use i2cdetect utility in format i2cdetect -y X, where X vary from 0 to 5 or more, to discover address occupied by device. To install utility you should run apt install i2c-tools on debian-kind system.

Getting help
------------

GoDoc [documentation](http://godoc.org/github.com/d2r2/go-hd44780)

Installation
------------

```bash
$ go get -u github.com/d2r2/go-bsbmp
```

License
-------

Go-dht is licensed under MIT License.
