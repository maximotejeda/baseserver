FROM golang:1.20 as build

WORKDIR /go/src/app

COPY . .

RUN go mod download && go mod tidy
RUN go build -o /go/bin/server ./

FROM debian:unstable-slim
COPY --from=build /go/bin/server /

ENTRYPOINT ["/server"]