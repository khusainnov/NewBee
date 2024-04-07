FROM golang:1.22 as build

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o newbee cmd/newbee/main.go

FROM gcr.io/distroless/base-debian12

COPY --from=build app/newbee .

COPY migrations ./migrations

EXPOSE 80

CMD ["/newbee"]
