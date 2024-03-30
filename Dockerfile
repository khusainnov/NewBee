FROM golang:1.22 as build

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o newbee cmd/newbee/main.go

FROM gcr.io/distroless/base-debian11

COPY --from=build app/newbee .


