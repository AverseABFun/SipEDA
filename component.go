package sipeda

type Reference string
type Library string
type Name string
type LibReference struct {
	Library Library
	Name    Name
}
type PinType uint8

const (
	PIN_Passive = PinType(iota)
	PIN_Input
	PIN_Output
	PIN_Bidirectional
	PIN_PowerInput
	PIN_PowerOutput
)

type Pin struct {
	Number string
	Name   string
	Type   PinType
}

type Component struct {
	LibRef            LibReference
	AllowedFootprints []LibReference
	DefaultFootprint  LibReference
	NumberOfPins      uint32
	Pins              []Pin
	RefTemplate       string
}
