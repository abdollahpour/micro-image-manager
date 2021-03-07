package processor

import (
	"strings"
)

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

// func (f *Format) String() ([]bimg.ImageType, error) {
// 	switch {
// 	case *f == "jpeg":
// 		return []bimg.ImageType{bimg.JPEG, bimg.WEBP}, nil
// 	case *f == "webp":
// 		return []bimg.ImageType{bimg.WEBP, bimg.JPEG}, nil
// 	case *f == "png":
// 		return []bimg.ImageType{bimg.PNG, bimg.WEBP}, nil
// 	default:
// 		return []bimg.ImageType{bimg.SVG, bimg.PNG, bimg.WEBP}, nil
// 	}
// }

type Profile struct {
	Name   string
	Width  int
	Height int
}

type ProcessingResult struct {
	File    string
	Profile Profile
	Format  Format
}

type ImagePocessor interface {
	Process(id string, bytes []byte, profiles []Profile) ([]ProcessingResult, error)
}
