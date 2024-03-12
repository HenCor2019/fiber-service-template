FROM golang:1.19-alpine AS base
ARG PORT

ENV GIT_VERSION 2.40.1-r0
ENV AIR_VERSION v1.40.2
ENV DIR $GOPATH/app/api

RUN apk add --no-cache git="$GIT_VERSION"

WORKDIR $DIR

FROM base AS dev

COPY go.* "$DIR"/

RUN go install github.com/cosmtrek/air@"$AIR_VERSION" && \
    go mod download

COPY . "$DIR"/

EXPOSE ${PORT}

CMD ["air", "-c", ".air.toml"]

FROM base AS build

COPY go.* "$DIR"/

RUN go mod download

COPY . "$DIR"/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s" \
    -o /go/bin/binary "$DIR"/cmd/app

FROM base AS staging

COPY --from=build /go/bin/binary /go/bin/binary

EXPOSE ${PORT}

CMD ["/go/bin/binary"]

FROM gcr.io/distroless/static-debian12:nonroot AS production
USER nonroot:nonroot

COPY internal "$DIR"/internal
COPY cmd "$DIR"/cmd

COPY --from=staging --chown=nonroot:nonroot /go/bin/binary /go/bin/binary

ENTRYPOINT ["/go/bin/binary"]
