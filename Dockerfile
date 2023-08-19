FROM golang:1.20 as base

WORKDIR /app
COPY . /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build cmd/rinha.go
FROM alpine

COPY go.mod go.sum ./
WORKDIR /app

COPY --from=base /app/rinha ./

EXPOSE 80
ENV GIN_MODE release
CMD [ "./rinha" ]