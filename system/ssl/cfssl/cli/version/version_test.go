package version

import (
	"testing"

	"github.com/ironstar-io/tokaido/system/ssl/cfssl/cli"
)

func TestVersionString(t *testing.T) {
	version := versionString()
	if version != "1.3.2" {
		t.Fatal("version string is not returned correctly")
	}
}

func TestVersionMain(t *testing.T) {
	args := []string{"cfssl", "version"}
	err := versionMain(args, cli.Config{})
	if err != nil {
		t.Fatal("version main failed")
	}
}
