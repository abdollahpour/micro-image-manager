package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/abdollahpour/micro-image-manager/internal/config"
	"github.com/abdollahpour/micro-image-manager/internal/processor"
	"github.com/abdollahpour/micro-image-manager/internal/server"
	"github.com/abdollahpour/micro-image-manager/internal/storage"
)

var Version = "development"

func main() {
	version := flag.Bool("version", false, "micro-image-manager version")

	flag.Parse()
	if *version {
		fmt.Println(Version)
		os.Exit(0)
	}

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

	imageProcessor := processor.NewBimgProcessor(conf.TempDir)
	imageStorage := storage.NewLocalStorage(conf.DistDir)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/images", server.StoreHandler(imageProcessor, imageStorage))
	mux.HandleFunc("/image", server.FetchHandler(imageStorage))
	log.Println(fmt.Sprintf("Listen on port %d", conf.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", conf.Host, conf.Port), mux))
}