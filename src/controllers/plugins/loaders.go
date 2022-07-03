package plugins

import (
	"log"
	"plugin"

	"github.com/FlamesX-128/anigo/src/types"
	"github.com/FlamesX-128/anigo/src/utils"
)

func putPlugin[T types.PluginType](p plugin.Symbol) {
	var u interface{} = p.(*types.Plugin[T]).This

	switch s := u.(type) {
	case types.Process:
		types.Processes = append(types.Processes, s)
	case types.Provider:
		types.Providers = append(types.Providers, s)
	case types.Service:
		types.Services = append(types.Services, s)
	}

}

func loader(path string) {
	plugin := utils.Try(plugin.Open(path))

	plug := utils.Try(plugin.Lookup("Plugin"))

	switch plug.(type) {
	case *types.Plugin[types.Process]:
		{
			putPlugin[types.Process](plug)

			break
		}
	case *types.Plugin[types.Provider]:
		{
			putPlugin[types.Provider](plug)

			break
		}
	case *types.Plugin[types.Service]:
		{
			putPlugin[types.Service](plug)

			break
		}
	default:
		{
			log.Panicln("Unknown plugin.")
			break
		}
	}

}

func multiloader(ps []string) {
	for _, p := range ps {
		loader(p)

	}

}

func Load(path string) {
	if ok, err := utils.Exists(path); !ok {
		if err == nil {
			log.Panicln("The directory " + path + " does not exist.")

		}

		return
	}

	multiloader(Search(path))
}
