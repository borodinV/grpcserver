FROM golang:1.20-alpine AS builder

RUN mkdir "/app"
COPY ./ ./app
WORKDIR ./app

RUN apk --no-cache add bash git make gcc gettext musl-dev
COPY ["go.mod","go.sum","./"]
RUN go mod download

RUN go build -o ./bin/grpcapp ./main.go

CMD ["./bin/grpcapp"]