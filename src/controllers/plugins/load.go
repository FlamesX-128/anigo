package plugins

import (
	"log"

	"github.com/FlamesX-128/anigo/src/utils"
)

func Load(path string) {
	if ok, err := utils.Exists(path); !ok {
		if err == nil {
			log.Panicln("The directory " + path + " does not exist.")

		}

		return
	}

	for _, p := range search(path) {
		if !loader(p) {
			log.Printf("An unknown plugin has been detected: %s\n", p)

		}

	}

}
