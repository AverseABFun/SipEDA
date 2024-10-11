package sipeda

import (
	"fmt"
	"io"
	"regexp"
	"strings"
)

type Line struct {
	Command  string
	Argument string
}

func readBytes(r io.Reader, numBytes uint) ([]byte, error) {
	var out = make([]byte, numBytes)
	var _, err = r.Read(out)
	return out, err
}

var re = regexp.MustCompile(`//.*`)

func SkipComment(r io.Reader) {
	var lastChar byte
	for lastChar != '\n' {
		var b, err = readBytes(r, 1)
		if err != nil {
			return
		}
		lastChar = b[0]
	}
}

func ParseLine(r io.Reader) (Line, error) {
	var in []byte
	var lastChar byte
	for lastChar != ';' {
		var b, err = readBytes(r, 1)
		if err != nil {
			if in != nil {
				fmt.Printf("Error: Before reaching ; got %v\n", err)
			}
			return Line{}, err
		}
		if lastChar == '/' && b[0] == '/' {
			SkipComment(r)
			return Line{}, nil
		}
		in = append(in, b[0])
		lastChar = b[0]
	}
	var ins = string(in)
	ins = re.ReplaceAllString(ins, "")
	ins = strings.TrimSpace(ins)
	if ins == "" {
		return Line{}, nil
	}
	ins = strings.TrimRight(ins, ";")

	var out = Line{}
	out.Command = strings.Split(ins, " ")[0]
	ins = strings.Replace(ins, out.Command, "", 1)
	ins = strings.TrimSpace(ins)
	out.Argument = ins

	return out, nil
}

func ParseFile(r io.Reader) ([]Line, error) {
	var err error = nil
	var out []Line
	for err == nil {
		var l Line
		l, err = ParseLine(r)
		if l.Command != "" {
			out = append(out, l)
		}
	}
	if err == io.EOF {
		return out, nil
	}
	return out, err
}
