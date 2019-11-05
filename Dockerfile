FROM golang:1.13 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /usr/bin/easycsr -v main.go

FROM scratch
COPY --from=builder /usr/bin/easycsr /easycsr

WORKDIR /csr
ENTRYPOINT [ "/easycsr" ]
