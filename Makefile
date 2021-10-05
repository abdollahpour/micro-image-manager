APP_VERSION:=edge
GOLANG_VERSION:=1.16
DOCKER_IMAGE:=abdollahpour/micro-image-manager
VIPS_VERSION:=8.10.0

compile:
	for i in darwin linux windows ; do \
		GOOS="$${i}" GOARCH=amd64 go build -o bin/mpg-"$${i}"-amd64 cmd/mpg/main.go; \
	done

archive:
	rm -f bin/*.zip
	for i in darwin linux windows ; do \
		zip -j "bin/mpg-$${i}-amd64.zip" "bin/mpg-$${i}-amd64" -x "*.DS_Store"; \
		zip "bin/mpg-$${i}-amd64.zip" -r templates; \
	done

run:
	go run cmd/server/main.go

get:
	go get -d -u ./...

image:
	docker build --pull \
		--cache-from "$(DOCKER_IMAGE):latest" \
		--build-arg GOLANG_VERSION="$(GOLANG_VERSION)" \
		--build-arg APP_VERSION="$(APP_VERSION)" \
		--tag "$(DOCKER_IMAGE):$(APP_VERSION)" \
		--file docker/Dockerfile .
	docker push "$(DOCKER_IMAGE):$(APP_VERSION)"
	# We update latest when a real version change happens
	if [ "$(APP_VERSION)" != "edge" ]; then \
		docker tag "$(DOCKER_IMAGE):$(APP_VERSION)" "$(DOCKER_IMAGE):latest"; \
		docker push "$(DOCKER_IMAGE):latest"; \
	fi

test_in_docker:
	docker build \
		--build-arg GOLANG_VERSION="$(GOLANG_VERSION)" \
		--tag test_in_docker \
		--file docker/Dockerfile-test .
	docker run -it -v $(shell pwd):/go/src/github.com/abdollahpour/micro-image-manager test_in_docker

test:
	go test -covermode=count -coverprofile=coverage.out -cover ./...

goveralls:
	$$GOPATH/bin/goveralls -service=travis-ci -coverprofile=coverage.out

pipeline:
	GITHUB_SHA=$$(git rev-parse HEAD)
	@for f in $(shell ls .tekton/pipeline-run-*); do cat $${f} | sed "s/latest/$${GITHUB_SHA}/g" | kubectl create -f -; done

dependencies:
	DEBIAN_FRONTEND=noninteractive VIPS_VERSION=8.10.0 \
	apt-get update && \
	apt-get install --no-install-recommends -y \
	ca-certificates \
	automake build-essential curl \
	gobject-introspection gtk-doc-tools libglib2.0-dev libjpeg62-turbo-dev libpng-dev \
	libwebp-dev libtiff5-dev libgif-dev libexif-dev libxml2-dev libpoppler-glib-dev \
	swig libmagickwand-dev libpango1.0-dev libmatio-dev libopenslide-dev libcfitsio-dev \
	libgsf-1-dev fftw3-dev liborc-0.4-dev librsvg2-dev libimagequant-dev libheif-dev && \
	cd /tmp && \
	curl -fsSLO https://github.com/libvips/libvips/releases/download/v${VIPS_VERSION}/vips-${VIPS_VERSION}.tar.gz && \
	tar zvxf vips-$(VIPS_VERSION).tar.gz && \
	cd /tmp/vips-$(VIPS_VERSION) && \
	CFLAGS="-g -O3" CXXFLAGS="-D_GLIBCXX_USE_CXX11_ABI=0 -g -O3" \
	./configure \
	--disable-debug \
	--disable-dependency-tracking \
	--disable-introspection \
	--disable-static \
	--enable-gtk-doc-html=no \
	--enable-gtk-doc=no \
	--enable-pyvips8=no && \
	make && \
	make install && \
	ldconfig
	go get ./...

build:
	CGO_ENABLED=0 GOOS=linux go build -ldflags="-X main.Version=1.0" -a -installsuffix cgo -o server ./cmd/server/main.go