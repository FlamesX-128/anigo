package utils

import (
	"log"
)

func Try[T any](data T, err error) T {
	if err != nil {
		log.Panicln(err)

	}

	return data
}
