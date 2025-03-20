# syntax=docker/dockerfile:1

FROM golang:1.23.2 AS build

WORKDIR $GOPATH/src/github.com/brotherlogic/githubstore

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 go build -o /githubstore

##
## Deploy
##
FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /githubstore /githubstore

EXPOSE 8080
EXPOSE 8081

USER nonroot:nonroot

ENTRYPOINT ["/githubstore"]