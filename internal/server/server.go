package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"

	"github.com/abdollahpour/micro-image-manager/internal/model"
	"github.com/abdollahpour/micro-image-manager/internal/processor"
	"github.com/abdollahpour/micro-image-manager/internal/storage"
	"github.com/google/jsonapi"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

func JSONError(w http.ResponseWriter, err interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}

type StoreHandlerResult struct {
	Id       string
	Profiles []model.Profile
	Formats  []model.Format
}

var (
	keyRe         = regexp.MustCompile(`profile_([a-z]+)`)
	valueRe       = regexp.MustCompile(`([0-9]{1,4})x([0-9]{1,4})`)
	profileNameRe = regexp.MustCompile(`[^a-z]`)
	imageRe       = regexp.MustCompile(`/image/([0-9a-zA-Z]{32}).([a-zA-Z]{3,4})`)
)

func DecodeProfile(key string, value string) (*model.Profile, error) {
	keyFounds := keyRe.FindStringSubmatch(key)
	if len(keyFounds) == 2 {
		valueFounds := valueRe.FindStringSubmatch(value)
		if len(valueFounds) == 3 {
			width, _ := strconv.Atoi(valueFounds[1])
			height, _ := strconv.Atoi(valueFounds[2])
			return &model.Profile{Name: keyFounds[1], Width: width, Height: height}, nil
		}
		return nil, errors.New("Value format is not currect")
	}

	return nil, nil
}

func StoreHandler(imageProcessor processor.ImagePocessor, imageStorage storage.Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		err := r.ParseMultipartForm(32 << 20) // 32Mb
		if err != nil {
			log.Warn("Request size if larger than 32Mb")
			w.WriteHeader(http.StatusBadRequest)
			jsonapi.MarshalErrors(w, []*jsonapi.ErrorObject{{
				Title:  "Multipart error",
				Detail: "Error to parse multipart POST request.",
				Status: "400",
				Code:   "REQ-100",
				Meta:   &map[string]interface{}{"spec": "https://tools.ietf.org/html/rfc2388"},
			}})
			return
		}

		var profiles []model.Profile
		for key, value := range r.Form {
			profile, err := DecodeProfile(key, value[0])
			if err != nil {
				log.WithField("key", key).WithField("value", value).Trace("Profile format error")
				w.WriteHeader(http.StatusBadRequest)
				jsonapi.MarshalErrors(w, []*jsonapi.ErrorObject{{
					Title:  "Format error",
					Detail: "Failed to parse profile.",
					Status: "400",
					Code:   "REQ-101",
					Meta:   &map[string]interface{}{"profile": key, "value": value[0]},
				}})
				return
			}
			if profile != nil {
				profiles = append(profiles, *profile)
			}
		}

		file, _, err := r.FormFile("image")
		if err != nil {
			log.Trace("No image found")
			w.WriteHeader(http.StatusBadRequest)
			jsonapi.MarshalErrors(w, []*jsonapi.ErrorObject{{
				Title:  "image not found",
				Detail: "No image file found in the POST request",
				Status: "400",
				Code:   "REQ-102",
			}})
			return
		}
		defer file.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			log.Error("Failed to read uploaded file")
			w.WriteHeader(http.StatusInternalServerError)
			jsonapi.MarshalErrors(w, []*jsonapi.ErrorObject{{
				Title:  "Internal server error",
				Status: "500",
				Code:   "INT-100",
			}})
			return
		}

		id := uuid.NewString()

		results, err := imageProcessor.Process(id, fileBytes, profiles)
		for _, result := range results {
			log.Error("Error to process the image")
			data, err := ioutil.ReadFile(result.File)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				jsonapi.MarshalErrors(w, []*jsonapi.ErrorObject{{
					Title:  "Internal server error",
					Status: "500",
					Code:   "INT-101",
				}})
				return
			}

			err = imageStorage.Store(id, result.Profile, result.Format, data)
			if err != nil {
				log.Error("Error to store the image")
				w.WriteHeader(http.StatusInternalServerError)
				jsonapi.MarshalErrors(w, []*jsonapi.ErrorObject{{
					Title:  "Internal server error",
					Status: "500",
					Code:   "INT-102",
				}})
				return
			}
		}

		result := StoreHandlerResult{
			Id:       id,
			Profiles: profiles,
			Formats:  []model.Format{model.JPEG, model.WEBP},
		}
		resultData, err := json.Marshal(result)
		if err != nil {
			log.Error("Error to serialize the result")
			w.WriteHeader(http.StatusInternalServerError)
			jsonapi.MarshalErrors(w, []*jsonapi.ErrorObject{{
				Title:  "Internal server error",
				Status: "500",
				Code:   "INT-103",
			}})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(resultData)
	}
}

func FetchHandler(imageStorage storage.Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		imageFounds := imageRe.FindStringSubmatch(r.URL.Path)
		if len(imageFounds) == 3 {
			id := imageFounds[1]
			format := imageFounds[2]
			profileName := profileNameRe.ReplaceAllString(r.URL.Query().Get("profile"), "")

			filePath, err := imageStorage.Fetch(id, model.NewProfile(profileName), model.NewFormat(format))
			if err != nil {
				fmt.Println(err)
			}

			w.Header().Set("Cache-Control", "public, max-age=604800, immutable")
			http.ServeFile(w, r, filePath)
			return
		}
		http.NotFoundHandler().ServeHTTP(w, r)
	}
}
