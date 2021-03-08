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
	go run cmd/mpg/main.go

image:
	docker build -t $(name) -f docker/Dockerfile . 

test:
	go test -coverprofile=coverage.out -cover ./...

coverage: test
	go tool cover -func coverage.out

spec:
	go test ./...

# Since solving it without put it in docker is complex because of dependencies, we do it here manually
coveralls:
	go test -v -covermode=count -coverprofile=coverage.out ./...
	$$HOME/go/bin/goveralls -coverprofile=coverage.out -service=travis-ci -repotoken $(TOKEN)
