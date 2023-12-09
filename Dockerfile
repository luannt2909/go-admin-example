# v0.0.1
FROM node:latest AS admin
RUN rm -rf ./admin/webs/.cache
RUN rm -rf ./admin/webs/dist
RUN rm -rf ./admin/webs/node_modules
COPY webs/package.json /admin/package.json
COPY webs/yarn.lock /admin/yarn.lock
WORKDIR /admin
RUN yarn install
ADD webs /admin
RUN WPATH='/admin' yarn run build

FROM --platform=linux/amd64 golang:alpine AS builder
RUN apk add --no-cache git gcc g++
#RUN apk add --no-cache git gcc g++ openssh-client build-base musl-dev
ADD . /go/src
WORKDIR /go/src

RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64  go build -ldflags="-w -s -extldflags=-static" -o go-admin .

FROM scratch
WORKDIR /root/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src .
COPY --from=admin /admin/dist ./admin/reminder-admin/dist/
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /usr/local/go/lib/time/

EXPOSE 8080
CMD [ "./go-admin" ]
