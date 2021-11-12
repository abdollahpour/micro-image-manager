APP_VERSION:=edge
GOLANG_VERSION:=1.16
DOCKER_IMAGE:=abdollahpour/micro-image-manager
VIPS_VERSION:=8.10.0

compile:
	for i in darwin linux windows ; do \
		GOOS="$${i}" GOARCH=amd64 go build -o bin/mim-"$${i}"-amd64 cmd/server/main.go; \
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
	go mod download

test_in_docker:
	docker build \
		--build-arg GOLANG_VERSION="$(GOLANG_VERSION)" \
		--tag test_in_docker \
		--file docker/Dockerfile-test .
	docker run -it -v $(shell pwd):/go/src/github.com/abdollahpour/micro-image-manager test_in_docker

test:
	go test -covermode=count -coverprofile=coverage.out -cover ./...
