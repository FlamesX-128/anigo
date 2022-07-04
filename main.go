package main

import (
	"sync"

	"github.com/FlamesX-128/anigo/src/types"
)

var wg sync.WaitGroup

func main() {
	aprops := types.Anigo_.Properties

	processes := aprops["persistent-processes"].([]types.ProcessPlugin[int32])
	fprocesses := aprops["floating-processes"].([]types.ProcessPlugin[float32])

	for _, process := range fprocesses {
		go process.Handler(&types.Anigo_)

	}

	for _, process := range processes {
		go func() {
			process.Handler(&types.Anigo_)
			wg.Done()
		}()

		wg.Add(1)
	}

	wg.Wait()
}
