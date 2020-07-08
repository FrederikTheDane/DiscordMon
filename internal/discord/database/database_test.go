package database

import (
	"os"
	"testing"
)

func TestConnect(t *testing.T) {
	if err := Connect(); err != nil {
		t.Fatal(err)
	}
}

func TestSetupTables(t *testing.T) {
	if err := Connect(); err != nil {
		t.Fatal(err)
	}

	if err := SetupTables(); err != nil {
		t.Fatal(err)
	}
}

func init() {
	os.Setenv("DISCORDMON_DB_HOST", "localhost")
	os.Setenv("DISCORDMON_DB_PORT", "5432")
	os.Setenv("DISCORDMON_DB_USER", "frederik")
	os.Setenv("DISCORDMON_DB_PWD", "frede")
	os.Setenv("DISCORDMON_DB_NAME", "discordmon")
}
