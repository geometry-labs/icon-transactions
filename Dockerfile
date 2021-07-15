FROM golang:1.16-buster AS builder

ARG SERVICE_NAME
ENV SERVICE_NAME ${SERVICE_NAME:-api}

# GO ENV VARS
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

# COPY SRC
WORKDIR /build
COPY ./src .

RUN go mod tidy

# CREATE SWAGGER DOCS
RUN go get github.com/swaggo/swag/cmd/swag
RUN go get github.com/alecthomas/template
RUN go get github.com/riferrei/srclient@v0.3.0
WORKDIR /build/api
RUN swag init -g routes/api.go

# BUILD
WORKDIR /build
RUN go build -o main ./${SERVICE_NAME}

FROM ubuntu as prod
COPY --from=builder /build/main /
CMD ["/main"]

FROM builder as test
CMD ["go", "test", "./.../", "-v", "-timeout", "15m"]
