package plugins

import (
	"log"

	"github.com/FlamesX-128/anigo/src/utils"
)

func Load(path string) bool {
	if ok, err := utils.Exists(path); !ok {
		if err == nil {
			log.Printf("The following directory doesn't exist: %s\n", path)

		}

		return false
	}

	for _, p := range search(path) {
		if !loader(p) {
			log.Printf("An unknown plugin has been detected: %s\n", p)

		}

	}

	return true
}
