package types

type PluginType interface {
	Process | Provider | Service
}

type Plugin[T PluginType] struct {
	Author, Name, Url, Version string
	This                       T
}

type Process struct {
	Handler    func(anigo *Anigo)
	Persistent bool
}

type Manager Process

type Provider struct {
	Handler func(string) (string, bool)
	Solvers []string
}

type Service struct {
	Handler map[string]func(...string) []any
	Solver  string
}
