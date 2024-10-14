package sipeda

type Format uint8

const (
	FORMAT_Invalid = Format(iota)
	FORMAT_Global
	FORMAT_Component
	FORMAT_Footprint
	FORMAT_PCB
	FORMAT_Schematic
	FORMAT_Project
)

func defaultFormatToString(f Format) string {
	switch f {
	case FORMAT_Global:
		return "FORMAT_Global"
	case FORMAT_Component:
		return "FORMAT_Component"
	case FORMAT_Footprint:
		return "FORMAT_Footprint"
	case FORMAT_PCB:
		return "FORMAT_PCB"
	case FORMAT_Schematic:
		return "FORMAT_Schematic"
	case FORMAT_Project:
		return "FORMAT_Project"
	default:
		return "FORMAT_Invalid"
	}
}

var FormatToString func(Format) string = defaultFormatToString

var Signatures = map[string]Format{
	"FORMAT_Component": FORMAT_Component,
	"FORMAT_Footprint": FORMAT_Footprint,
	"FORMAT_PCB":       FORMAT_PCB,
	"FORMAT_Schematic": FORMAT_Schematic,
	"FORMAT_Project":   FORMAT_Project,
}

var validFORMAT_FootprintCommands = []string{
	"FORMAT_Footprint",
	"endFORMAT_Footprint",
	"lib",
	"name",
	"pads",
	"units",
	"font",
	"layer",
	"pad",
	"draw",
	"x",
	"y",
	"w",
	"h",
	"d",
	"rect",
	"circle",
	"rotate",
	"enddraw",
	"deselect",
	"align",
	"text",
	"from",
	"to",
	"copy",
}

var ValidCommands = map[Format][]string{
	FORMAT_Global: {
		"version",
		"type",
	},
	FORMAT_Component: {
		"add_fp_lib",
		"FORMAT_Component",
		"lib",
		"name",
		"pins",
		"ref",
		"FORMAT_Footprint",
		"default_FORMAT_Footprint",
		"pin",
		"type",
		"name",
		"endFORMAT_Component",
	},
	FORMAT_Footprint: {
		"FORMAT_Footprint",
		"endFORMAT_Footprint",
		"lib",
		"name",
		"pads",
		"units",
		"font",
		"layer",
		"pad",
		"draw",
		"x",
		"y",
		"w",
		"h",
		"d",
		"points",
		"rect",
		"circle",
		"rotate",
		"enddraw",
		"deselect",
		"align",
		"text",
		"from",
		"to",
		"copy",
	},
	FORMAT_PCB: {
		"FORMAT_Footprint",
		"lib",
		"name",
		"pads",
		"units",
		"font",
		"layer",
		"pad",
		"draw",
		"x",
		"y",
		"w",
		"h",
		"d",
		"points",
		"rect",
		"circle",
		"rotate",
		"enddraw",
		"deselect",
		"align",
		"text",
		"from",
		"to",
		"copy",
		"add_sch",
		"add_fp_lib",
		"p",
		"track",
		"ref",
	},
	FORMAT_Schematic: {
		"add_cmp_lib",
		"rail",
		"FORMAT_Component",
		"ref",
		"FORMAT_Footprint",
		"desc",
		"point",
		"connect",
	},
	FORMAT_Project: {
		"FORMAT_Schematic",
		"FORMAT_PCB",
	},
}
