package main

import (
	"github.com/volatiletech/sqlboiler/v4/drivers"
	"gotest/sqlboiler-vertica/driver"
)

func main() {
	drivers.DriverMain(&driver.VerticaDriver{})
}
