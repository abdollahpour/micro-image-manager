package server

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"runtime"
	"strings"
	"testing"

	"github.com/abdollahpour/micro-image-manager/internal/model"
	"github.com/abdollahpour/micro-image-manager/internal/processor"
	"github.com/abdollahpour/micro-image-manager/internal/storage"
	"github.com/stretchr/testify/assert"
)

type mockImageProcessor struct{}

func (m *mockImageProcessor) Process(id string, bytes []byte, profiles []model.Profile) ([]model.ProcessingResult, error) {
	return nil, nil
}

type mockImageStorage struct{}

func (m *mockImageStorage) Store(id string, profile model.Profile, format model.Format, data []byte) error {
	return nil
}

func (m *mockImageStorage) Fetch(id string, profile model.Profile, format model.Format) (string, error) {
	return "", nil
}

func TestStoreHandler(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	filePath := strings.Replace(filename, ".go", ".jpeg", 1)

	file, err := os.Open(filePath)
	assert.Nil(t, err)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", "sample.jpg")
	io.Copy(part, file)

	profileLarge := model.Profile{
		Name:   "large",
		Width:  800,
		Height: 600,
	}
	profileSmall := model.Profile{
		Name:   "small",
		Width:  400,
		Height: 300,
	}

	writer.WriteField("profile_large", "800x600")
	writer.WriteField("profile_small", "400x300")

	writer.Close()

	req, err := http.NewRequest("POST", "/api/v1/images", body)
	assert.Nil(t, err)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	imageProcessor := processor.NewBimgProcessor(os.TempDir())
	imageStorage := storage.NewLocalStorage(os.TempDir())

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(StoreHandler(imageProcessor, imageStorage))

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	// assert.Equal(t, "application/json", rr.HeaderMap["Content-Type"][0])
	// assert.Equal(t, "public, max-age=604800, immutable", rr.HeaderMap["Cache-Control"][0])

	var result StoreHandlerResult
	json.Unmarshal(rr.Body.Bytes(), &result)

	assert.ElementsMatch(t, []model.Format{model.JPEG, model.WEBP}, result.Formats)
	assert.ElementsMatch(t, []model.Profile{profileLarge, profileSmall}, result.Profiles)
}

func TestStoreHandlerPostNotMultipart(t *testing.T) {
	req, err := http.NewRequest("POST", "/api/v1/images", strings.NewReader("data=value"))
	assert.Nil(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(StoreHandler(&mockImageProcessor{}, &mockImageStorage{}))

	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Result().StatusCode)
	assert.Equal(t, "{\"errors\":[{\"title\":\"Multipart error\",\"detail\":\"Error to parse multipart POST request.\",\"status\":\"400\",\"code\":\"REQ-100\",\"meta\":{\"spec\":\"https://tools.ietf.org/html/rfc2388\"}}]}\n", rr.Body.String())
}

func TestDecodeProfile(t *testing.T) {
	profile, err := DecodeProfile("somekey", "somevalue")
	assert.Nil(t, profile)
	assert.Nil(t, err)

	profile, err = DecodeProfile("profile_large", "invalidvalue")
	assert.Nil(t, profile)
	assert.NotNil(t, err)

	profile, err = DecodeProfile("profile_large", "800x600")
	assert.Equal(t, &model.Profile{Name: "large", Width: 800, Height: 600}, profile)
	assert.Nil(t, err)
}
