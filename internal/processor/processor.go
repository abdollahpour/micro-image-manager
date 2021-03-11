package processor

import "github.com/abdollahpour/micro-image-manager/internal/model"

type ImagePocessor interface {
	Process(id string, bytes []byte, profiles []model.Profile) ([]model.ProcessingResult, error)
}
