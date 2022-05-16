FROM golang:1.18.2-alpine3.14 as builder

RUN apk add --no-cache make

WORKDIR /tflint
COPY . /tflint
RUN make build

FROM alpine:3.14.2 as prod

LABEL maintainer=terraform-linters

RUN apk add --no-cache ca-certificates

COPY --from=builder /tflint/dist/tflint /usr/local/bin

ENTRYPOINT ["tflint"]
WORKDIR /data
