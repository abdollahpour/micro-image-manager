FROM abdollahpour/micro-image-manager:builder as builder
COPY internal internal/
COPY cmd cmd/
RUN GOOS=linux go build \
  -a -installsuffix cgo \
  -ldflags="-X 'main.Version=${APP_VERSION}'" \
  -o /usr/local/bin/mim-server cmd/server/main.go
CMD sh

FROM debian:buster-slim

COPY --from=builder /usr/local/lib /usr/local/lib
COPY --from=builder /usr/local/bin/mim-server /usr/local/bin/mim-server
COPY --from=builder /etc/ssl/certs /etc/ssl/certs

# Install runtime dependencies
RUN DEBIAN_FRONTEND=noninteractive \
  apt-get update && \
  apt-get install --no-install-recommends -y \
  procps libglib2.0-0 libjpeg62-turbo libpng16-16 libopenexr23 \
  libwebp6 libwebpmux3 libwebpdemux2 libtiff5 libgif7 libexif12 libxml2 libpoppler-glib8 \
  libmagickwand-6.q16-6 libpango1.0-0 libmatio4 libopenslide0 \
  libgsf-1-114 fftw3 liborc-0.4-0 librsvg2-2 libcfitsio7 libimagequant0 libheif1 && \
  apt-get autoremove -y && \
  apt-get autoclean && \
  apt-get clean && \
  rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*


# Server port to listen
ENV PORT 8080
ENV MIM_DIST_DIR /mnt/images

RUN mkdir -p /mnt/images

# Run the entrypoint command by default when the container starts.
ENTRYPOINT ["/usr/local/bin/mim-server"]

# Expose the server TCP port
EXPOSE ${PORT}