package processor

import (
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

func toTargetImageTypes(format model.Format) ([]bimg.ImageType, error) {
	switch {
	case format == model.JPEG:
		return []bimg.ImageType{bimg.JPEG, bimg.WEBP}, nil
	case format == model.PNG:
		return []bimg.ImageType{bimg.PNG, bimg.WEBP}, nil
	case format == model.WEBP:
		return []bimg.ImageType{bimg.WEBP, bimg.JPEG}, nil
	case format == model.SVG:
		return []bimg.ImageType{bimg.SVG, bimg.WEBP, bimg.PNG}, nil
	default:
		return nil, fmt.Errorf("Format not supported: %v", format)
	}
}

func toFormat(bimgFormat bimg.ImageType) (model.Format, error) {
	switch {
	case bimgFormat == bimg.JPEG:
		return model.JPEG, nil
	case bimgFormat == bimg.PNG:
		return model.PNG, nil
	case bimgFormat == bimg.WEBP:
		return model.WEBP, nil
	case bimgFormat == bimg.SVG:
		return model.SVG, nil
	default:
		return model.NOT_SUPPORTED, fmt.Errorf("Format not supported: %v", bimgFormat)
	}
}

func (p BimgProcessor) Process(id string, bytes []byte, profiles []model.Profile) ([]model.ProcessingResult, error) {
	image := bimg.NewImage(bytes)
	imageType := image.Type()

	format, err := model.NewFormat(imageType)
	if err != nil {
		return nil, err
	}

	imageTypes, err := toTargetImageTypes(*format)
	if err != nil {
		return nil, err
	}

	results := make([]model.ProcessingResult, len(profiles)*len(imageTypes))

	for i, profile := range profiles {
		resized, err := image.ResizeAndCrop(profile.Width, profile.Height)
		if err != nil {
			return nil, err
		}

		for j, imageType := range imageTypes {
			converted, err := bimg.NewImage(resized).Convert(imageType)
			if err != nil {
				return nil, err
			}

			path := filepath.Join(p.tempDir, fmt.Sprintf("%s_%s.%v", id, profile.Name, bimg.ImageTypeName(imageType)))
			bimg.Write(path, converted)

			format, err := toFormat(imageType)
			if err != nil {
				return nil, err
			}

			results[i*len(imageTypes)+j] = model.ProcessingResult{
				File:    path,
				Profile: profile,
				Format:  format,
			}
		}

		// Reset tht image. In an unexpected bahaviar bimg also change size of the source image
		if i < len(imageTypes)-1 {
			image = bimg.NewImage(bytes)
		}
	}

	return results, nil
}
