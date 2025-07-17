FROM golang:1.24.0-alpine AS builder
WORKDIR /app
COPY go.mod .
COPY main.go .

RUN <<EOF
go mod tidy 
go build
EOF

FROM scratch
WORKDIR /app
COPY --from=builder /app/simplest-agent .

ENTRYPOINT ["./simplest-agent"]
