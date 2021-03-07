package storage

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
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

func (s *LocalStorage) Store(id string, profileName string, format string, data []byte) error {
	filePath := path.Join(s.distDir, fmt.Sprintf("%s_%s.%s", id, profileName, format))
	return ioutil.WriteFile(filePath, data, 0644)
}

func (s *LocalStorage) Fetch(id string, profileName string, format string) (string, error) {
	path := path.Join(s.distDir, fmt.Sprintf("%s_%s.%s", id, profileName, format))
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return "", errors.New(path + " not found")
	}
	return path, nil
}
