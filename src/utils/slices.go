package utils

import "github.com/FlamesX-128/anigo/src/types"

func Contains[T types.SliceContent](sc []T, c T) bool {
	for _, e := range sc {
		if c == e {

			return true
		}

	}

	return false
}

func Filter[T any](sc []T, test func(T) bool) (r []T) {
	for _, s := range sc {
		if test(s) {
			r = append(r, s)

		}

	}

	return
}
