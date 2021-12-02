#  BASE IMAGE
FROM golang:1.17-alpine as base
RUN apk update && apk upgrade && \
	apk add --update alpine-sdk && \
	apk add --no-cache musl-dev git make cmake build-base
ENV GIT_TERMINAL_PROMPT=1


#  DEV IMAGE
FROM base as dev
# ADD in air file monitor
RUN go install github.com/cosmtrek/air@latest
COPY ./docker/app/air.conf /tmp/air.conf
WORKDIR /go/reunion
# SETUP LIVE RELOAD
CMD air -c /tmp/air.conf


#  PRODUCTION BINARY IMAGE BUILD
FROM node:16-alpine as builder
COPY --from=golang:1.17-alpine /usr/local/go/ /usr/local/go/
ENV PATH="/usr/local/go/bin:${PATH}"
ADD . /go/reunion
WORKDIR /go/reunion/ui
RUN npm rebuild esbuild
RUN npm install
RUN npm run build
RUN ls dist

# get modules before build
WORKDIR /go/reunion
ENV GO111MODULE on
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download -x

# build with build info
RUN go build -i -v /go/reunion


# final no frills binary
FROM iron/base as binary
# COPY IN COMPILED BINARY
WORKDIR /app
ENV PATH=$PATH:/app
COPY --from=builder /go/reunion/reunion /app/reunion
COPY --from=builder /go/reunion/ui/dist /app/ui/dist
COPY --from=builder /go/reunion/.env /app/.env
ENTRYPOINT ["reunion"]
