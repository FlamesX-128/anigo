package types

type SliceContent interface {
	complex64 | complex128 | float32 | float64 |
		uint8 | uint16 | uint32 | uint | uintptr |
		int8 | int16 | int32 | int |
		bool | string
}
