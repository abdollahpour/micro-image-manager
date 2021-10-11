FROM abdollahpour/go-libvip:1.17 as builder
ARG APP_VERSION=edge

WORKDIR ${GOPATH}/src/github.com/abdollahpour/micro-image-manager

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY internal internal/
COPY cmd cmd/

RUN GOOS=linux go build \
  -a -installsuffix cgo \
  -ldflags="-X 'main.Version=${APP_VERSION}'" \
  -o /usr/local/bin/mim-server cmd/server/main.go
CMD sh

FROM abdollahpour/go-libvip:1.17

COPY --from=builder /usr/local/bin/mim-server /usr/local/bin/mim-server

# Server port to listen
ENV PORT 8080
ENV MIM_DIST_DIR /mnt/images

RUN mkdir -p /mnt/images

# Run the entrypoint command by default when the container starts.
ENTRYPOINT ["/usr/local/bin/mim-server"]

# Expose the server TCP port
EXPOSE ${PORT}