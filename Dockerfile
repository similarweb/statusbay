##########################################
# Build golang container
##########################################
FROM golang:1.13.4-alpine AS go-builder

# Install OS level dependencies
RUN apk add --update alpine-sdk git && \
	git config --global http.https://gopkg.in.followRedirects true 

# Set workdir for the rest of the commands
WORKDIR /app

# Now add project files
COPY . .

# Build a binary
RUN go install -v ./...

##########################################
# Runtime container
##########################################
FROM alpine:3.8
RUN apk add --update ca-certificates python py-pip && rm -rf /var/cache/apk/* && \
	apk add curl jq && \
	pip install awscli && \
    apk --purge -v del py-pip && \
	wget https://releases.hashicorp.com/consul-template/0.22.0/consul-template_0.22.0_linux_386.zip  -O consul-template.zip && \
	unzip consul-template.zip -d /usr/local/bin/

COPY --from=go-builder /go/bin/statusbay /bin/statusbay

COPY start.sh .
EXPOSE 8080
ENTRYPOINT [ "sh", "start.sh" ]