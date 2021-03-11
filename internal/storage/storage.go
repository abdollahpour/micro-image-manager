package storage

import "github.com/abdollahpour/micro-image-manager/internal/model"

// Storage images
type Storage interface {
	Store(id string, profile model.Profile, format model.Format, data []byte) error
	Fetch(id string, profile model.Profile, format model.Format) (string, error)
}
