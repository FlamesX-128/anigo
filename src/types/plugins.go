package types

type ProcessPlugin[_ float32 | int32] struct {
	Handler func(*Anigo)
	Test    func() []Status

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
	ProcessPlugin[int32] |
		ProcessPlugin[float32] |
		ProviderPlugin |
		ServicePlugin
}
