FROM docker.io/golang:1.23.2-alpine3.20 AS builder
WORKDIR /src

COPY ./go.mod ./go.sum ./
RUN go mod download
WORKDIR /src/api
COPY ./api ./

RUN go build -ldflags="-s -w" -o /web

# Final stage, ref https://github.com/GoogleContainerTools/distroless/blob/main/base/README.md for distroless
FROM gcr.io/distroless/static
COPY --from=builder /web .
USER 1000
EXPOSE 8001
ENTRYPOINT ["/web"]
