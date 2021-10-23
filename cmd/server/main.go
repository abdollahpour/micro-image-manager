package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/abdollahpour/micro-image-manager/internal/config"
	"github.com/abdollahpour/micro-image-manager/internal/processor"
	"github.com/abdollahpour/micro-image-manager/internal/server"
	"github.com/abdollahpour/micro-image-manager/internal/storage"
	log "github.com/sirupsen/logrus"
)

var Version = "development"

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	version := flag.Bool("version", false, "print version version")
	debug := flag.Bool("debug", false, "active debug manager")

	flag.Parse()
	if *version {
		fmt.Println(Version)
		os.Exit(0)
	}
	if *debug {
		log.SetLevel(log.TraceLevel)
	}

	conf := config.NewEnvConfiguration()

	if _, err := os.Stat(conf.TempDir); os.IsNotExist(err) {
		err = os.Mkdir(conf.TempDir, 0744)
		if err != nil {
			log.Fatal("Failed to create temp dir: " + conf.TempDir)
		}
	}

	imageProcessor := processor.NewBimgProcessor(conf.TempDir)
	imageStorage := storage.NewLocalStorage(conf.DistDir)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/images", server.StoreHandler(imageProcessor, imageStorage))
	mux.HandleFunc("/image/", server.FetchHandler(imageStorage))
	mux.HandleFunc("/live", server.Live(imageStorage))
	log.Println(fmt.Sprintf("Listen on port %d", conf.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", conf.Host, conf.Port), mux))
}
