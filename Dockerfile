# App builder
FROM golang:1.20-bookworm AS builder

WORKDIR /go/src/github.com/kamaal111/pocket-slate-api/
COPY . .
# Download dependencies.
RUN apt update && apt install -y \
    tzdata ca-certificates
RUN go mod download -x
RUN go mod verify
# Run update certificates
RUN update-ca-certificates
# Build the binary.
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -ldflags="-w -s" -v -o /go/bin/pocket-slate-api src/*.go

# Build a smaller image with the minimum required things to run.
FROM scratch
# Import from builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/pocket-slate-api /go/bin/pocket-slate-api
# Run the pocket-slate-api binary.
EXPOSE 8000
ENTRYPOINT ["/go/bin/pocket-slate-api"]
