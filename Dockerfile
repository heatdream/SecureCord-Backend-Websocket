FROM golang:latest as websocket_builder

WORKDIR /srv/application/
ADD . /srv/application/
RUN go mod download && go mod verify
ENV GOOS=linux GOARCH=amd64
RUN go build -o websocket-exec .


FROM alpine:latest
WORKDIR /root/
COPY --from=websocket_builder /srv/application/websocket-exec ./
CMD ["./websocket-exec"]
