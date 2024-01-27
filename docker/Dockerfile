### Build Project

FROM golang:1.21.6 as builder

WORKDIR /healthproxy/

COPY .. .

RUN go mod download
RUN go mod verify
RUN CGO_ENABLED=0 go build -o healthproxy

### Make Image

FROM alpine:latest

WORKDIR /healthproxy

COPY --from=builder /healthproxy/healthproxy /healthproxy/
COPY docker/*.sh .
RUN ls

CMD [ "/bin/sh", "-c", "/healthproxy/entrypoint.sh" ]