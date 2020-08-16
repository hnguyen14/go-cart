FROM golang:1.14.7 AS build
ARG REVISION

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN make

FROM alpine:3.12
ARG REVISION

WORKDIR /app

COPY --from=build /app/dist/go-cart ./
COPY ./public ./public
COPY ./templates ./templates
COPY ./config.yml ./

RUN adduser -u 1000 gocart -D && \
    chown -R gocart .
USER 1000

CMD ["./go-cart"]
