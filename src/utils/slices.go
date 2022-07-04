package utils

import "github.com/FlamesX-128/anigo/src/types"

func Contains[T types.SliceContent](ss []T, s T) bool {
	for _, e := range ss {
		if e == s {

			return true
		}

	}

	return false
}

func Filter[T interface{}](ss []T, test func(T) bool) (r []T) {
	for _, s := range ss {
		if test(s) {
			r = append(r, s)

		}

	}

	return
}
