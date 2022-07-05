package utils

import (
	"log"

	"github.com/FlamesX-128/anigo/src/types"
)

func GetKeys[T types.SliceContent, Y any](m map[T]Y) (k []T) {
	for e := range m {
		k = append(k, e)

	}

	return
}

func Try[T any](data T, err error) T {
	if err != nil {
		log.Panicln(err)

	}

	return data
}
