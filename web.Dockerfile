FROM --platform=$BUILDPLATFORM docker.io/golang:1.23.2-alpine3.20 AS builder
ARG TARGETARCH
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=${TARGETARCH}

WORKDIR /src

COPY ./go.mod ./go.sum ./
RUN go mod download

WORKDIR /src/api
COPY ./web ./

RUN go build -ldflags="-s -w" -o /web

# Final stage, ref https://github.com/GoogleContainerTools/distroless/blob/main/base/README.md for distroless
FROM gcr.io/distroless/static
COPY --from=builder /web .
USER 1000
EXPOSE 8000
ENTRYPOINT ["/web"]
