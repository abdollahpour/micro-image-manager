APP_VERSION:=edge
GOLANG_VERSION:=1.16
DOCKER_IMAGE:=abdollahpour/micro-image-manager

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