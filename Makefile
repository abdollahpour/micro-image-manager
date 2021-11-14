APP_VERSION:=edge
GOLANG_VERSION:=1.16
DOCKER_IMAGE:=abdollahpour/micro-image-manager
VIPS_VERSION:=8.10.0

compile_in_docker:
	for i in darwin linux windows ; do \
		GOOS=linux go build \
		-a -installsuffix cgo \
		-ldflags="-X 'main.Version=$${APP_VERSION}'" \
		-o /release/mim-"$${i}"-amd64 cmd/server/main.go; \
    done

archive_in_docker:
	for i in darwin linux windows ; do \
    	zip -j "/release/mim-$${i}-amd64.zip" "/release/mim-$${i}-amd64" -x "*.DS_Store"; \
    done

release:
	mkdir -p bin
	rm -rf bin/*
	docker run -v $(shell pwd)/bin:/release --rm $$(docker build --no-cache --build-arg APP_VERSION=$(VIPS_VERSION) -q -f docker/Dockerfile-release .)

run:
	go run cmd/server/main.go

get:
	go mod download

test_in_docker:
	docker build \
		--build-arg GOLANG_VERSION="$(GOLANG_VERSION)" \
		--tag test_in_docker \
		--file docker/Dockerfile-test .
	docker run -it -v $(shell pwd):/go/src/github.com/abdollahpour/micro-image-manager test_in_docker

test:
	go test -covermode=count -coverprofile=coverage.out -cover ./...
