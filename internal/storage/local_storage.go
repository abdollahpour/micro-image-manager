package storage

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/abdollahpour/micro-image-manager/internal/model"
)

// LocalStorage uses local directory to store files
type LocalStorage struct {
	distDir string
}

// NewLocalStorage create a new localStorage
func NewLocalStorage(distDir string) *LocalStorage {
	return &LocalStorage{
		distDir: distDir,
	}
}

func (s *LocalStorage) Store(id string, profile model.Profile, format model.Format, data []byte) error {
	var profileName string
	if profile.Default {
		profileName = model.DefaultProfile.Name
	} else {
		profileName = profile.Name
	}
	filePath := path.Join(s.distDir, fmt.Sprintf("%s_%s.%s", id, profileName, format))
	return ioutil.WriteFile(filePath, data, 0644)
}

func (s *LocalStorage) Fetch(id string, profile model.Profile, format model.Format) (string, error) {
	var filePath string
	filePath = path.Join(s.distDir, fmt.Sprintf("%s_%v.%s", id, profile.Name, format))
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		filePath = path.Join(s.distDir, fmt.Sprintf("%s_%v.%s", id, model.DefaultProfile.Name, format))
		_, err := os.Stat(filePath)

		if os.IsNotExist(err) {
			return "", errors.New(filePath + " not found")
		}
	}
	return filePath, nil
}
