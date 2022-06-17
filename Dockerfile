FROM golang:1.16 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

ENV CGO_ENABLED=0
RUN go build -ldflags='-w -s -extldflags "-static"' -a -v /app

FROM scratch
COPY --from=builder /app/server /

CMD [ "/server" ]