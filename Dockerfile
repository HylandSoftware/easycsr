FROM golang:1.11 AS builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 make build-unix

FROM scratch
COPY --from=builder /app/dist/easycsr /easycsr

WORKDIR /csr
ENTRYPOINT [ "/easycsr" ]