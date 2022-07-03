package main

import (
	"os"
	"path"
	"sync"

	"github.com/FlamesX-128/anigo/src/controllers/plugins"
	"github.com/FlamesX-128/anigo/src/types"
	"github.com/FlamesX-128/anigo/src/utils"
)

var Scope = &types.Anigo{}

func main() {
	var processes []func(anigo *types.Anigo)
	var wg sync.WaitGroup

	dir := utils.Try(os.Getwd())

	plugins.Load(path.Join(dir, "plugins"))

	for _, p := range types.Processes {
		if p.Persistent {
			processes = append(processes, p.Handler)

			continue
		}

		go p.Handler(Scope)
	}

	for _, p := range processes {
		wg.Add(1)

		go (func() {
			p(Scope)
			defer wg.Done()
		})()

	}

	wg.Wait()
}
