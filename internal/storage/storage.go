package storage

// Storage images
type Storage interface {
	Store(id string, profileName string, format string, data []byte) error
	Fetch(id string, profileName string, format string) (string, error)
}
