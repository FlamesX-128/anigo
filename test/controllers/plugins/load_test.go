package plugins

import (
	"os"
	"path"
	"testing"

	"github.com/FlamesX-128/anigo/src/controllers/plugins"
	"github.com/FlamesX-128/anigo/src/utils"
)

func Test_0(t *testing.T) {
	directory := path.Join(utils.Try(os.Getwd()), "..", "..", "temp")

	if !plugins.Load(directory) {
		t.Errorf("Expected: %v, got: %v", true, false)

	}

	if plugins.Load("321-plugins") {
		t.Errorf("Expected: %v, got: %v", false, true)

	}

}
