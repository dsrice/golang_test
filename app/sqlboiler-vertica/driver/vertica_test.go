// These tests assume there is a user sqlboiler_driver_user and a database
// by the name of sqlboiler_driver_test that it has full R/W rights to.
// In order to create this you can use the following steps from a root
// crdb account:
//
//   create role sqlboiler_driver_user login nocreatedb nocreaterole nocreateuser password 'sqlboiler';
//   create database sqlboiler_driver_test owner = sqlboiler_driver_user;

package driver

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"testing"

	"github.com/volatiletech/sqlboiler/v4/drivers"
)

var (
	flagOverwriteGolden = flag.Bool("overwrite-golden", false, "Overwrite the golden file with the current execution results")

	envHostname = drivers.DefaultEnv("DRIVER_HOSTNAME", "localhost")
	envPort     = drivers.DefaultEnv("DRIVER_PORT", "26257")
	envUsername = drivers.DefaultEnv("DRIVER_USER", "root")
	envPassword = drivers.DefaultEnv("DRIVER_PASS", "")
	envDatabase = drivers.DefaultEnv("DRIVER_DB", "sqlboiler_driver_test")
)

func TestAssemble(t *testing.T) {

	config := drivers.Config{
		"sslmode": "disable",
		"schema":  "public",
	}

	c := &VerticaDBDriver{}
	info, err := c.Assemble(config)
	if err != nil {
		t.Fatal(err)
	}

	got, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	if *flagOverwriteGolden {
		if err = ioutil.WriteFile("vertica.golden.json", got, 0664); err != nil {
			t.Fatal(err)
		}
		t.Log("wrote:", string(got))
		return
	}

	want, err := ioutil.ReadFile("vertica.golden.json")
	if err != nil {
		t.Fatal(err)
	}

	if bytes.Compare(want, got) != 0 {
		t.Errorf("want:\n%s\ngot:\n%s\n", want, got)
	}
}
