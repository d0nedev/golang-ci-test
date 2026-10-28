FROM golang:1.27 AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /app .

FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /app
COPY --from=build /app /app/app
COPY config.yaml /app/config.yaml
EXPOSE 8080
ENTRYPOINT ["/app/app"]
