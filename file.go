package sipeda

import "errors"

func GetVersionFromFile(ls []Line) (string, error) {
	for _, l := range ls {
		if l.Command == "version" {
			return l.Argument, nil
		}
	}
	return "", errors.New("got no signature")
}

func GetFormatFromFile(ls []Line) (Format, error) {
	for _, l := range ls {
		if l.Command == "type" {
			var f, ok = Signatures[l.Argument]
			if !ok {
				return FORMAT_Invalid, errors.New("got unknown signature")
			}
			return f, nil
		}
	}
	return FORMAT_Invalid, errors.New("got no signature")
}
