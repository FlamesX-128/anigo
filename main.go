package main

import (
	"os"
	"path"
	"sync"

	"github.com/FlamesX-128/anigo/src/controllers/plugins"
	"github.com/FlamesX-128/anigo/src/types"
	"github.com/FlamesX-128/anigo/src/utils"
)

var wg sync.WaitGroup

func main() {
	plugins.Load(
		path.Join(utils.Try(os.Getwd()), "plugins"),
	)

	var (
		pprocesses = types.Anigo_.Properties["persistent-processes"]
		fprocesses = types.Anigo_.Properties["floating-processes"]
	)

	for _, process := range fprocesses.([]types.ProcessPlugin[float32]) {
		go process.Handler(&types.Anigo_)

	}

	for _, process := range pprocesses.([]types.ProcessPlugin[int32]) {
		go func() {
			process.Handler(&types.Anigo_)
			wg.Done()
		}()

		wg.Add(1)
	}

	wg.Wait()
}
