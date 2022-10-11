package main

import (
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/drivers"
	"gotest/sqlboiler-vertica/driver"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "Version: v4")
		return
	}
	drivers.DriverMain(&driver.VerticaDBDriver{})
}
