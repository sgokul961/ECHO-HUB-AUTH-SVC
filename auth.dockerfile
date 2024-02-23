FROM golang:1.22.0-alpine3.19 as build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY cmd cmd
COPY pkg pkg

RUN go build -o /api -v ./cmd
FROM alpine3.19
COPY --from=build /api /api


EXPOSE 3000


CMD ["/api"]