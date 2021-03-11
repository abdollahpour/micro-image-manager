package model

import "strings"

type Format string

const (
	JPEG Format = "JPEG"
	GIF         = "GIF"
	PNG         = "PNG"
	WEBP        = "WEBP"
	SVG         = "SVG"
)

var formatToString = map[Format]string{
	JPEG: "JPEG",
	GIF:  "GIF",
	PNG:  "PNG",
	WEBP: "WEBP",
	SVG:  "SVG",
}

var stringToFormat = map[string]Format{
	"JPEG": JPEG,
	"GIF":  GIF,
	"PNG":  PNG,
	"WEBP": WEBP,
	"SVG":  SVG,
}

func NewFormat(value string) Format {
	return stringToFormat[strings.ToUpper(value)]
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
