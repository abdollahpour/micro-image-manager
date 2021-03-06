package storage

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/abdollahpour/micro-image-manager/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestLocalStorageStoreAndFetch(t *testing.T) {
	localStorage := NewLocalStorage(os.TempDir())

	id := "279fc0dd-4160-4a10-ac43-4702477049ec"
	data := []byte{1, 2, 3}
	err := localStorage.Store(id, model.Profile{Name: "large"}, "jpeg", data)
	assert.Nil(t, err)
	fetched, err := localStorage.Fetch(id, model.Profile{Name: "large"}, "jpeg")
	defer os.Remove(fetched)
	assert.Nil(t, err)
	fetchedData, err := ioutil.ReadFile(fetched)
	assert.Nil(t, err)
	assert.Equal(t, data, fetchedData)
}

func TestLocalStorageFileNotExist(t *testing.T) {
	localStorage := NewLocalStorage(os.TempDir())

	id := "279fc0dd-4160-4a10-ac43-4702477049ec"
	_, err := localStorage.Fetch(id, model.NewProfile("large"), "jpeg")
	assert.NotNil(t, err)
}
func TestLocalStorgeStoreDefaultProfile(t *testing.T) {
	localStorage := NewLocalStorage(os.TempDir())

	id := "279fc0dd-4160-4a10-ac43-4702477049ec"

	smallData := []byte{1, 2, 3}
	err := localStorage.Store(id, model.DefaultProfile, "jpeg", smallData)
	assert.Nil(t, err)
	assertFetchedData(t, *localStorage, id, "small", smallData)

	largeData := []byte{1, 2, 3, 4}
	err = localStorage.Store(id, model.NewProfile("large"), "jpeg", largeData)
	assert.Nil(t, err)
	assertFetchedData(t, *localStorage, id, "large", largeData)

	err = localStorage.Store(id, model.DefaultProfile, "jpeg", largeData)
	assert.Nil(t, err)
	assertFetchedData(t, *localStorage, id, "non_existing_profile", largeData)
}

func assertFetchedData(t *testing.T, localStorage LocalStorage, id string, profileName string, data []byte) {
	fetched, err := localStorage.Fetch(id, model.NewProfile(profileName), "jpeg")
	defer os.Remove(fetched)
	assert.Nil(t, err)
	fetchedData, err := ioutil.ReadFile(fetched)
	assert.Nil(t, err)
	assert.Equal(t, data, fetchedData)
}
