FROM alpine:latest

# Update package index and install go + git
RUN apk add --update go git
RUN apk add --no-cache musl-dev

# Set up GOPATH
RUN mkdir /go
ENV GOPATH /go

# Get dependencies
RUN go get github.com/deckarep/golang-set

# Add source code
ADD . /go/src/github.com/b6luong/Snakes-and-Ladders

# Build and Install
RUN go install github.com/b6luong/Snakes-and-Ladders/cmd/slgame

# Where to run the application
ENTRYPOINT /go/bin/slgame
