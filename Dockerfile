# Build Stage
FROM --platform=linux/amd64 golang:latest as builder

## Install build dependencies.
RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive apt-get install -y unzip

## Add source code to the build stage.
ADD . /rare
WORKDIR /rare

## TODO: ADD YOUR BUILD INSTRUCTIONS HERE.
RUN go mod download
RUN go build .
RUN ls /rare

FROM --platform=linux/amd64 golang:latest

## TODO: Change <Path in Builder Stage>
COPY --from=builder /rare/rare /rare
RUN ls /
