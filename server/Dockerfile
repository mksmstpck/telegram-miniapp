ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm as builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -C cmd -v -o /cmd/cmd .


FROM debian:bookworm

COPY --from=builder /cmd/cmd /usr/local/bin/
CMD ["cmd"]
