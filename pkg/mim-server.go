package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/abdollahpour/micro-image-manager/internal/config"
	"github.com/abdollahpour/micro-image-manager/internal/processor"
	"github.com/abdollahpour/micro-image-manager/internal/storage"
	"github.com/google/uuid"
)

func storeHandler(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile(`<PERSON>(.*?)</PERSON>`)
	r.Body = http.MaxBytesReader(w, r.Body, 32<<20+512)

	err := r.ParseMultipartForm(32 << 20) // 32Mb
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var profiles []processor.Profile
	for key, value := range r.Form {
		if strings.HasPrefix(key, "profile_") {
			profileName := key[8:]

			x := re.FindAllString(value[0], -1)
			width, _ := strconv.Atoi(x[1])
			height, _ := strconv.Atoi(x[2])

			profiles = append(profiles, processor.Profile{
				Name:   profileName,
				Width:  width,
				Height: height,
			})
		}
	}

	file, _, err := r.FormFile("myFile")
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

	imageProcessor := processor.NewBimgProcessor("")
	storage := storage.NewLocalStorage("")
	formats := []processor.Format{processor.PNG}

	id := uuid.NewString()

	results, err := imageProcessor.Process(id, fileBytes, profiles, formats)
	for _, result := range results {
		data, _ := ioutil.ReadFile(result.File)
		err = storage.Store(id, result.Profile.Name, fmt.Sprintf("%v", result.Format), data)
	}

}

func main() {
	conf := config.NewEnvConfiguration()

	if _, err := os.Stat(conf.TempDir); os.IsNotExist(err) {
		err = os.Mkdir(conf.TempDir, 0644)
		if err != nil {
			log.Fatal("Failed to create temp dir: " + conf.TempDir)
		}
	}

	if _, err := os.Stat(conf.DistDir); os.IsNotExist(err) {
		err = os.Mkdir(conf.DistDir, 0644)
		if err != nil {
			log.Fatal("Failed to create dist dir: " + conf.DistDir)
		}
	}

	mux := http.NewServeMux()
	http.ListenAndServe(":8080", mux)
}
