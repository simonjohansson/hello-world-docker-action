FROM golang:1.23-bookworm AS base

WORKDIR $GOPATH/src/hello-world-docker-action

COPY . .

RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/static-debian12

COPY --from=base /go/bin/app /

CMD ["/app"]
