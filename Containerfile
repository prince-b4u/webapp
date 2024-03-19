# Use a minimal base image
FROM docker.io/golang:1.22.1  AS builder

ENV CGO_ENABLED=0
ENV GOOS=linux

WORKDIR /app

COPY go.mod .

RUN go mod tidy && go mod download && go mod verify

COPY . .

RUN go build -v -ldflags "-w -s" -o /usr/local/bin/app ./cmd/webapp/webapp.go

# Create a minimal final image
FROM alpine

ENV URL="/webapp"

COPY --from=builder /usr/local/bin/app /usr/local/bin/app

EXPOSE 3000
CMD ["app"]
