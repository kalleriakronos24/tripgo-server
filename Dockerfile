FROM golang:latest as base

RUN apt-get update && apt-get install -y \
    git \
    curl \
    docker.io

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

RUN apt-get clean && rm -rf /var/lib/apt/lists/*

RUN go env -w GO111MODULE=auto

COPY . /opt/app/api

WORKDIR /opt/app/api

RUN go mod tidy && go get ./ && go build -buildvcs=auto && rm -rf .git

CMD ["air"]


FROM nginx:latest

COPY /nginx/default.conf /etc/nginx/nginx.conf

CMD ["nginx", "-g", "daemon off"]