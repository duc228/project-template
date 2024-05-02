FROM golang:1.22.2

WORKDIR /opt

RUN go install github.com/cosmtrek/air@v1.51.0

# air -c .air.toml
CMD tail -f /dev/null