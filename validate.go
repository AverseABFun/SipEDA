package sipeda

import (
	"fmt"
	"slices"
)

func ValidateFile(ls []Line) error {
	var ver, err = GetVersionFromFile(ls)
	if err != nil {
		return err
	}
	var ok = CheckFileVersion(ver)
	if ok.Error != nil {
		return ok.Error
	}
	format, err := GetFormatFromFile(ls)
	if err != nil {
		return err
	}
	for _, l := range ls {
		if !slices.Contains(ValidCommands[format], l.Command) && !slices.Contains(ValidCommands[FORMAT_Global], l.Command) {
			return fmt.Errorf("invalid command %s", l.Command)
		}
	}
	return nil
}
