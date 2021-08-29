FROM golang:1.16-alpine AS build
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
ADD . /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o /graphql-cardkeeper-server .

FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=build /graphql-cardkeeper-server /graphql-cardkeeper-server
EXPOSE 8080
USER nonroot:nonroot

ENTRYPOINT ["/graphql-cardkeeper-server", "--config", "config.toml"]