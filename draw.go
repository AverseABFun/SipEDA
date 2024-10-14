package sipeda

import "golang.org/x/image/font/sfnt"

type Position struct {
	X float64
	Y float64
}
type Size Position
type Rotation float64

type Rect struct {
	Pos  Position
	Size Size
}

type Circle struct {
	Pos    Position
	Radius float64
}

type Align uint8

const (
	ALIGN_Left = Align(iota)
	ALIGN_Center
	ALIGN_Right
)

type Text struct {
	Pos   Position
	Font  *sfnt.Font
	Size  float64
	Align Align
	Text  string
}
