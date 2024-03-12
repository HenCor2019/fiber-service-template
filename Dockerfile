FROM golang:1.19-alpine AS base
ARG PORT

ENV GIT_VERSION 2.40.1-r0
ENV DIR $GOPATH/app/api

RUN apk add --no-cache git="$GIT_VERSION"

WORKDIR $DIR

FROM base AS dev

COPY go.mod $DIR
COPY go.sum $DIR

RUN go install github.com/cosmtrek/air@v1.40.2 && \
    go mod download

COPY . "$DIR"/

EXPOSE ${PORT}

CMD ["air", "-d"]

FROM base AS build

COPY go.mod "$DIR"
COPY go.sum "$DIR"
RUN go mod download

COPY . "$DIR"/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s" \
    -o /go/bin/binary "$DIR"/cmd/app

FROM base AS production

COPY --from=build /go/bin/binary "$DIR"/main

EXPOSE ${PORT}

CMD ["./main"]

FROM gcr.io/distroless/static-debian12:nonroot AS final
USER nonroot:nonroot

COPY . "$DIR"/

COPY --from=build --chown=nonroot:nonroot /go/bin/binary /go/bin/binary

ENTRYPOINT ["/go/bin/binary"]
