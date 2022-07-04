package types

import (
	"log"
	"reflect"
	"strings"
)

// //---------------------------------------------------------// //

func _GetServicesOrProviders[T ServicePlugin | ProviderPlugin](
	a Anigo, property string) ([]T, bool) {

	if prop, ok := a.Properties[property]; ok {
		if prop, ok := prop.([]T); ok {

			return prop, true
		}

		log.Printf("The \"%s\" property is unknown.\n", property)
	}

	return []T{}, false

}

// //---------------------------------------------------------// //

func _GetServicesOrProvidersNames[T ServicePlugin | ProviderPlugin](
	a Anigo, prop string) ([]string, bool) {

	if props, ok := _GetServicesOrProviders[T](a, prop); ok {

		var names []string

		for _, prop := range props {

			solvers := reflect.ValueOf(prop).
				FieldByName("Solvers").
				Interface()

			names = append(names, solvers.([]string)...)

		}

		return names, true

	}

	return []string{}, false

}

// //---------------------------------------------------------// //

func _GetDomain(url string) string {
	u := strings.Replace(url, "https://", "", 1)

	return u[:strings.Index(u, "/")]
}
