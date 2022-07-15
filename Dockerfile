ARG NAME=cards
ARG SOURCEROOT=/go/src/${NAME}

# Builder Image
FROM golang:1.18-alpine as builder

ARG NAME
ARG SOURCEROOT

COPY . ${SOURCEROOT}
WORKDIR ${SOURCEROOT}

RUN apk add build-base
RUN go mod vendor -v
RUN GOOS=linux go build -o bin/${NAME} cmd/main.go

# Runner Image
FROM alpine:latest
ARG NAME
ARG SOURCEROOT
WORKDIR /usr/bin

RUN apk update && apk add bash && apk --no-cache add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=builder ${SOURCEROOT}/.env.example /usr/bin/.env
COPY --from=builder ${SOURCEROOT}/bin/${NAME} /usr/bin

CMD ["cards"]
