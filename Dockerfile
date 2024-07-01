##########################################
# Build golang container
##########################################
FROM golang:1.22.4-alpine AS go-builder
ARG GOFLAGS=""
ARG GOCACHE="/go/.cache/go-build"
ENV GOFLAGS="$GOFLAGS"

# Install OS level dependencies
RUN apk add --update alpine-sdk git && \
	git config --global http.https://gopkg.in.followRedirects true

# Set workdir for the rest of the commands
WORKDIR /app

# Create an cache layer for dependencies
COPY go.mod go.sum ./
RUN go env -w GOMODCACHE=$GOCACHE; go env -w GOCACHE=$GOCACHE
RUN --mount=type=cache,target=$GOCACHE go mod download

# Build the binary
COPY . .
RUN --mount=type=cache,target=$GOCACHE go build -o /bin/statusbay

##########################################
# Runtime container
##########################################
FROM alpine:3.20 AS runtime
COPY --from=go-builder /bin/statusbay /bin/statusbay

RUN mkdir -p /etc/statusbay/
COPY events.yaml /etc/statusbay/
EXPOSE 8080
ENTRYPOINT ["/bin/statusbay"]
