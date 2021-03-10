package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"

	"github.com/abdollahpour/micro-image-manager/internal/processor"
	"github.com/abdollahpour/micro-image-manager/internal/storage"
	"github.com/google/uuid"
)

func JSONError(w http.ResponseWriter, err interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}

type StoreHandlerResult struct {
	Id       string
	Profiles []processor.Profile
	Formats  []processor.Format
}

var (
	keyRe   = regexp.MustCompile(`profile_([a-z]+)`)
	valueRe = regexp.MustCompile(`([0-9]{1,4})x([0-9]{1,4})`)
	imageRe = regexp.MustCompile(`/image/([0-9a-zA-Z]{32}).([a-zA-Z]{3,4})`)
)

func DecodeProfile(key string, value string) (*processor.Profile, error) {
	keyFounds := keyRe.FindStringSubmatch(key)
	if len(keyFounds) == 2 {
		valueFounds := valueRe.FindStringSubmatch(value)
		if len(valueFounds) == 3 {
			width, _ := strconv.Atoi(valueFounds[1])
			height, _ := strconv.Atoi(valueFounds[2])
			return &processor.Profile{Name: keyFounds[1], Width: width, Height: height}, nil
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
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var profiles []processor.Profile
		for key, value := range r.Form {
			profile, err := DecodeProfile(key, value[0])
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if profile != nil {
				profiles = append(profiles, *profile)
			}
		}

		file, _, err := r.FormFile("image")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}
		defer file.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}

		id := uuid.NewString()

		results, err := imageProcessor.Process(id, fileBytes, profiles)
		for _, result := range results {
			data, _ := ioutil.ReadFile(result.File)
			err = imageStorage.Store(id, result.Profile.Name, result.Format.String(), data)
			if err != nil {
				fmt.Println(err)
			}
		}

		result := StoreHandlerResult{
			Id:       id,
			Profiles: profiles,
			Formats:  []processor.Format{processor.JPEG, processor.WEBP},
		}
		resultData, _ := json.Marshal(result)

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
			profile := r.URL.Query().Get("profile")

			filePath, err := imageStorage.Fetch(id, profile, format)
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
