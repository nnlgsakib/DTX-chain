FROM golang:1.18-alpine AS builder

WORKDIR /dtx-edge

ADD go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o dtx-edge main.go

FROM alpine:latest AS runner

RUN apk --no-cache add ca-certificates jq

WORKDIR /dtx-edge

COPY --from=builder /dtx-edge/dtx-edge ./
COPY ./docker/local/dtx-edge.sh ./

# Expose json-rpc, libp2p and grpc ports
EXPOSE 8545 9632 1478

ENTRYPOINT ["./dtx-edge.sh"]
