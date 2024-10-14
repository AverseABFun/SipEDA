package sipeda

import (
	_ "embed"

	"golang.org/x/image/font/sfnt"
	"golang.org/x/image/math/fixed"
)

//go:embed OpenSans-Regular.ttf
var openSans []byte
var OpenSans *sfnt.Font

func PopulateFontData() error {
	sans, err := sfnt.Parse(openSans)
	if err != nil {
		return err
	}
	OpenSans = sans
	return nil
}

// ppem is number of pixels in one em
func GetSegmentsForGlyph(f sfnt.Font, ppem fixed.Int26_6, r rune) (sfnt.Segments, error) {
	var idx, err = f.GlyphIndex(nil, r)
	if err != nil {
		return nil, err
	}
	return f.LoadGlyph(nil, idx, ppem, nil)
}
