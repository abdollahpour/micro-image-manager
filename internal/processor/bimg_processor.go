package processor

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/abdollahpour/micro-image-manager/internal/model"
	"github.com/h2non/bimg"
)

type BimgProcessor struct {
	tempDir string
}

func NewBimgProcessor(tempDir string) *BimgProcessor {
	return &BimgProcessor{
		tempDir: tempDir,
	}
}

func sourceTypeToTargetFormats(imageType string) ([]bimg.ImageType, error) {
	switch {
	case imageType == "jpeg":
		return []bimg.ImageType{bimg.JPEG, bimg.WEBP}, nil
	case imageType == "png":
		return []bimg.ImageType{bimg.PNG, bimg.WEBP}, nil
	case imageType == "webp":
		return []bimg.ImageType{bimg.WEBP, bimg.JPEG}, nil
	case imageType == "svg":
		return []bimg.ImageType{bimg.SVG, bimg.WEBP, bimg.PNG}, nil
	default:
		return nil, errors.New("Format not supported: " + imageType)
	}
}

func (p BimgProcessor) Process(id string, bytes []byte, profiles []model.Profile) ([]model.ProcessingResult, error) {
	image := bimg.NewImage(bytes)
	imageType := image.Type()
	formats, err := sourceTypeToTargetFormats(imageType)
	if err != nil {
		return nil, err
	}

	results := make([]model.ProcessingResult, len(profiles)*len(formats))

	for i, profile := range profiles {
		resized, err := image.Resize(profile.Width, profile.Height)
		if err != nil {
			return nil, err
		}

		for j, format := range formats {
			converted, err := bimg.NewImage(resized).Convert(format)
			if err != nil {
				return nil, err
			}

			path := filepath.Join(p.tempDir, fmt.Sprintf("%s_%s.%v", id, profile.Name, bimg.ImageTypeName(format)))
			bimg.Write(path, converted)

			results[i*len(formats)+j] = model.ProcessingResult{
				File:    path,
				Profile: profile,
				Format:  model.NewFormat(bimg.ImageTypeName(format)),
			}
		}
	}

	return results, nil
}
