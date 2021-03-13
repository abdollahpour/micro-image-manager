package model

import "strings"

type Format string

const (
	JPEG Format = "jpeg"
	GIF         = "gif"
	PNG         = "png"
	WEBP        = "webp"
	SVG         = "svg"
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

func NewFormat(value string) Format {
	return stringToFormat[strings.ToLower(value)]
}

func (f *Format) String() string {
	return formatToString[*f]
}

type Profile struct {
	Name    string
	Width   int
	Height  int
	Default bool
}

var DefaultProfile = Profile{Name: "default", Default: true}

func NewProfile(name string) Profile {
	return Profile{Name: name, Default: len(name) == 0}
}

type ProcessingResult struct {
	File    string
	Profile Profile
	Format  Format
}
