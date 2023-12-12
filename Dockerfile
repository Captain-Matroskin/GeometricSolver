# Builder
FROM golang:1.20.12-alpine3.18 AS builderGeomSolver
WORKDIR /cont
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN apk update && apk upgrade && \
    apk --update add git make
RUN go build -o geomSolver ./cmd/geomSolver/main.go

FROM alpine:latest
RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app
WORKDIR /app

COPY --from=builderGeomSolver ./cont/geomSolver /app

CMD ["/app/geomSolver"]