package plugins

import (
	"log"
	"plugin"

	"github.com/FlamesX-128/anigo/src/types"
)

func loader(path string) bool {
	plugin, err := plugin.Open(path)

	if err != nil {
		log.Println(err.Error())

		return false
	}

	plug, _ := plugin.Lookup("Plugin")

	switch data := plug.(type) {
	case types.ProcessPlugin[int32]:
		{
			types.Anigo_.Properties["persistent-processes"] = append(
				types.Anigo_.Properties["persistent-processes"].([]types.ProcessPlugin[int32]), data,
			)
			break
		}

	case types.ProcessPlugin[float32]:
		{
			types.Anigo_.Properties["floating-processes"] = append(
				types.Anigo_.Properties["floating-processes"].([]types.ProcessPlugin[float32]), data,
			)
			break
		}

	case types.ProviderPlugin:
		{
			types.Anigo_.Properties["providers"] = append(
				types.Anigo_.Properties["providers"].([]types.ProviderPlugin), data,
			)
			break
		}

	case types.ServicePlugin:
		{
			types.Anigo_.Properties["services"] = append(
				types.Anigo_.Properties["services"].([]types.ServicePlugin), data,
			)
			break
		}
	default:
		return false

	}

	return true
}
