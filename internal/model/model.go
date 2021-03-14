package model

import (
	"errors"
	"sort"
	"strings"
)

type Format string

const (
	JPEG          Format = "jpeg"
	GIF                  = "gif"
	PNG                  = "png"
	WEBP                 = "webp"
	SVG                  = "svg"
	NOT_SUPPORTED        = "not_supported"
)

var formatToString = map[Format]string{
	JPEG: "jpeg",
	GIF:  "gif",
	PNG:  "png",
	WEBP: "webp",
	SVG:  "svg",
}

var stringToFormat = map[string]Format{
	"jpeg": JPEG,
	"gif":  GIF,
	"png":  PNG,
	"webp": WEBP,
	"svg":  SVG,
}

func NewFormat(value string) (*Format, error) {
	format, exist := stringToFormat[strings.ToLower(value)]
	if exist {
		return &format, nil
	}
	return nil, errors.New("foramt not found for value: " + value)
}

func (f *Format) String() string {
	return formatToString[*f]
}

type Profile struct {
	Name   string
	Width  int
	Height int
}

var DefaultProfile = Profile{Name: "default"}

func NewProfile(name string) Profile {
	return Profile{Name: name}
}

type ProcessingResult struct {
	File    string
	Profile Profile
	Format  Format
}

func SortProfile(profiles []Profile) {
	sort.SliceStable(profiles, func(i, j int) bool {
		return profiles[i].Width*profiles[i].Height > profiles[j].Width*profiles[j].Height
	})
}
