package types

type PersistentProcess int32
type FloatingProcess float32

type ProcessPlugin[_ ~float32 | ~int32] struct {
	Handler func(*Anigo)
	Test    []Status

	Name string
}

type ProviderPlugin struct {
	Handler func(string, string) (string, bool)
	Test    func() []Status

	Solvers []string
	Name    string
}

type ServicePlugin struct {
	Handler func(string, ...string) (interface{}, bool)
	Test    func() []Status

	Solvers []string
	Name    string
}

type PluginType interface {
	ProcessPlugin[PersistentProcess] |
		ProcessPlugin[FloatingProcess] |
		ProviderPlugin |
		ServicePlugin
}
