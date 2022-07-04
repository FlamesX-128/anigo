package types

type Anigo struct {
	Properties map[string]interface{}
}

type Status struct {
	Status string
	Name   string
}

type SliceContent interface {
	complex64 | complex128 | float32 | float64 |
		uint8 | uint16 | uint32 | uint | uintptr |
		int8 | int16 | int32 | int |
		bool | string
}

var Anigo_ = Anigo{
	Properties: map[string]interface{}{
		"persistent-processes": []ProcessPlugin[int32]{},
		"floating-processes":   []ProcessPlugin[float32]{},

		"providers": []ProviderPlugin{},
		"services":  []ProviderPlugin{},
	},
}
