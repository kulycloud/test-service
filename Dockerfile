FROM golang:1.15.3-alpine AS builder

ADD test-service/go.mod test-service/go.sum /build/test-service/
ADD protocol/go.mod protocol/go.sum /build/protocol/
ADD common/go.mod common/go.sum /build/common/

ENV CGO_ENABLED=0

WORKDIR /build/test-service
RUN go mod download

COPY test-service/ /build/test-service/
COPY protocol/ /build/protocol
COPY common/ /build/common
RUN go build -o /build/kuly .

FROM scratch

COPY --from=builder /build/kuly /

CMD ["/kuly"]
