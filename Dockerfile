FROM golang:alpine as build

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o gola .

FROM scratch


WORKDIR /app

EXPOSE ":9000"

COPY --from=build /app/gola .

COPY .env /app/

ENTRYPOINT ["/app/gola"]


