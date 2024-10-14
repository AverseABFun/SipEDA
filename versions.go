package sipeda

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type CommitHash string

type Version struct {
	Major       uint16
	Minor       uint16
	Patch       uint16
	Development bool
	Commit      CommitHash
}

func (v Version) GetFullVersionString() string {
	var str = "release"
	if v.Development {
		str = "dev"
	}
	var str2 = ""
	if v.Commit != "" {
		str2 = "-" + string(v.Commit)
	}
	return fmt.Sprintf("%d.%d.%d-%s%s", v.Major, v.Minor, v.Patch, str, str2)
}

func (v Version) GetFileVersion() string {
	return fmt.Sprintf("%d.%d", v.Major, v.Minor)
}

func GetVersionFromString(str string) (Version, error) {
	var segments = strings.Split(str, "-")
	if str == "" {
		return Version{}, errors.New("received empty string")
	}
	var segments2 = strings.Split(segments[0], ".")
	var out = Version{}
	if len(segments2) < 1 {
		return out, errors.New("no major version")
	}
	var major, err = strconv.Atoi(segments2[0])
	if err != nil {
		return out, err
	}
	out.Major = uint16(major)
	if len(segments2) < 2 {
		return out, errors.New("no minor version")
	}
	minor, err := strconv.Atoi(segments2[1])
	if err != nil {
		return out, err
	}
	out.Minor = uint16(minor)
	if len(segments2) < 3 {
		return out, errors.New("no patch version")
	}
	patch, err := strconv.Atoi(segments2[2])
	if err != nil {
		return out, err
	}
	out.Patch = uint16(patch)
	if len(segments) < 2 {
		return out, errors.New("no development/release information")
	}
	out.Development = segments[1] == "dev"
	if len(segments) < 3 {
		return out, errors.New("no commit hash")
	}
	out.Commit = CommitHash(segments[2])
	return out, nil
}

var CURRENT_VERSION = Version{Major: 0, Minor: 1, Patch: 0, Development: true, Commit: "00fae1a"} //set commit after every commit!!
var CURRENT_FILE_VERSION = CURRENT_VERSION.GetFileVersion()

type ErrWarn struct {
	Error    error
	Warnings []error
}

func (ew *ErrWarn) AddWarning(warning error) {
	ew.Warnings = append(ew.Warnings, warning)
}

func CheckFileVersion(str string) ErrWarn {
	var ver, err = GetVersionFromString(str)
	if ver.Major == 0 && ver.Minor == 0 {
		return ErrWarn{Error: err}
	}
	if ver.Major < CURRENT_VERSION.Major {
		return ErrWarn{Error: errors.New("too low major version")}
	}
	var out = ErrWarn{}
	if ver.Major > CURRENT_VERSION.Major {
		out.AddWarning(errors.New("higher major version"))
	}
	if ver.Minor < CURRENT_VERSION.Minor && ver.Major == CURRENT_VERSION.Major {
		out.AddWarning(errors.New("lower minor version"))
	}
	return out
}
