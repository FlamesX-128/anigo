package utils

import (
	"github.com/FlamesX-128/anigo/src/types"
)

func GetKeys[T types.SliceContent, Y any](m map[T]Y) (k []T) {
	for e := range m {
		k = append(k, e)

	}

	return
}
