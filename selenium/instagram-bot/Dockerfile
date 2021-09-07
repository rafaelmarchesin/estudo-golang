FROM golang:alpine3.12 AS base
RUN echo 'http://dl-cdn.alpinelinux.org/alpine/v3.6/main' >> /etc/apk/repositories
RUN echo 'http://dl-cdn.alpinelinux.org/alpine/v3.6/community' >> /etc/apk/repositories
RUN apk add --update --no-cache git ca-certificates perl openssh-client
RUN apk add --update --no-cache mysql
RUN go get github.com/codegangsta/gin
RUN go get golang.org/x/tools/cmd/godoc
RUN go get github.com/axw/gocov/gocov
RUN mkdir /db-test
ENV CGO_ENABLED 0
WORKDIR /app

# Save cache to speed up changes on go modules
COPY build-input/go.mod go.mod
COPY build-input/go.sum go.sum
RUN go mod download

# Update go modules to latest version
COPY app/go.mod go.mod
COPY app/go.sum go.sum
RUN go mod download

# Copy test scripts to image
COPY scripts/* /usr/local/bin/
RUN chmod a+x /usr/local/bin/*

# Install MySQL client for testing
RUN apk add mysql-client

# Install go-swagger
RUN apk add --update curl jq
RUN download_url=$(curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/latest | \
  jq -r '.assets[] | select(.name | contains("'"$(uname | tr '[:upper:]' '[:lower:]')"'_amd64")) | .browser_download_url') \
  && curl -o /usr/local/bin/swagger -L'#' "$download_url"
RUN chmod +x /usr/local/bin/swagger

FROM base AS dev
ENTRYPOINT sh -c "echo godoc at: http://$(show-ip|tr -d ' '):6060; (godoc -http 0.0.0.0:6060 > /dev/null &); echo application at: http://$(show-ip|tr -d ' '):3000; echo generating swagger.json; swagger generate spec --scan-models -o swagger.json; gin --appPort 8080 --bin build/gin-bin run main.go"

FROM base AS build
COPY app/ /app
RUN validate-gofmt

# Generate swagger.json
RUN swagger generate spec --scan-models -o swagger.json

RUN CI=true validate-test-coverage
RUN go build -o app .

FROM alpine AS prod
WORKDIR /app
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/app /app/app
COPY --from=build /app/swagger.json /app/swagger.json
ENTRYPOINT sh -c "./app"