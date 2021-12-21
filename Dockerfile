FROM quay.io/iamplus/golang:1.14 AS builder

WORKDIR /microservice

COPY ./ ./

RUN make tools
#RUN make lint
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags maf -ldflags '-w -extldflags "-static"' -o microserviceentry


FROM quay.io/iamplus/alpine:3.8

RUN apk --no-cache add ca-certificates tzdata shadow && groupadd -r nonroot && useradd --no-log-init -rm -g nonroot nonroot

USER nonroot

WORKDIR /microservice

COPY --from=builder  /microservice/microserviceentry .
ENTRYPOINT ["./microserviceentry"]  