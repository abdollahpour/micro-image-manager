package processor

import (
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBimgProcessor(t *testing.T) {
	bimgProcessor := NewBimgProcessor(os.TempDir())

	_, filename, _, _ := runtime.Caller(0)
	filePath := strings.Replace(filename, ".go", ".png", 1)
	fileData, err := ioutil.ReadFile(filePath)
	assert.Nil(t, err)

	id := "279fc0dd-4160-4a10-ac43-4702477049ec"
	large := Profile{Name: "large", Width: 800, Height: 600}
	small := Profile{Name: "small", Width: 400, Height: 300}
	profiles := []Profile{large, small}

	results, err := bimgProcessor.Process(id, fileData, profiles)
	assert.Nil(t, err)
	assert.Equal(t, []ProcessingResult{
		{
			File:    path.Join(os.TempDir(), id+"_large.png"),
			Profile: large,
			Format:  PNG,
		},
		{
			File:    path.Join(os.TempDir(), id+"_large.webp"),
			Profile: large,
			Format:  WEBP,
		},
		{
			File:    path.Join(os.TempDir(), id+"_small.png"),
			Profile: small,
			Format:  PNG,
		},
		{
			File:    path.Join(os.TempDir(), id+"_small.webp"),
			Profile: small,
			Format:  WEBP,
		},
	}, results)
}