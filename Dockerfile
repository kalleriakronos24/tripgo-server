FROM golang:latest AS base
COPY . /opt/app/api
WORKDIR /opt/app/api
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

FROM base AS init
RUN apt-get update && apt-get install -y \
    git \
    curl \
    docker.io
RUN apt-get clean && rm -rf /var/lib/apt/lists/*
RUN go env -w GO111MODULE=auto

FROM base AS go-builder-production
COPY . .
RUN go mod tidy && go get ./
RUN go build -buildvcs=false -v -o ./bin/server ./main.go
#RUN go build -ldflags="-s -w" -o ./bin/server ./main.go

FROM base AS go-builder-developement
RUN go mod tidy && go get ./


FROM alpine:latest AS production
WORKDIR /usr/bin
COPY --from=go-builder-production /opt/app/api/bin /usr/bin
CMD ["./usr/bin/server"]

FROM go-builder-developement AS development
COPY --from=go-builder-developement . /opt/app/api
CMD ["air"]
