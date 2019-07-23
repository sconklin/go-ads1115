package main

import (
	"github.com/sconklin/go-ads1115"
	"github.com/sconklin/go-i2c"
	logger "github.com/sconklin/go-logger"
	"time"
)

var lg = logger.NewPackageLogger("main",
	logger.DebugLevel,
	// logger.InfoLevel,
)

func main() {
	defer logger.FinalizeLogger()
	// Create new connection to i2c-bus on 1 line with address 0x76.
	// Use i2cdetect utility to find device address over the i2c-bus
	i2c, err := i2c.NewI2C(0x48, 1)
	if err != nil {
		lg.Fatal(err)
	}
	defer i2c.Close()

	lg.Notify("***************************************************************************************************")
	lg.Notify("*** You can change verbosity of output, to modify logging level of modules \"i2c\", \"ads1115\"")
	lg.Notify("*** Uncomment/comment corresponding lines with call to ChangePackageLogLevel(...)")
	lg.Notify("***************************************************************************************************")
	// Uncomment/comment next lines to suppress/increase verbosity of output
	logger.ChangePackageLogLevel("i2c", logger.InfoLevel)
	logger.ChangePackageLogLevel("ads", logger.InfoLevel)

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
