FROM golang:alpine as source
WORKDIR /home/server
COPY . .
WORKDIR cmd/compression-service
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -mod vendor -o compression-service

FROM alpine as runner
LABEL name="bitstored/compression-service"
RUN apk --update add ca-certificates
COPY --from=source /home/server/cmd/compression-service/compression-service /home/compression-service
COPY --from=source /home/server/scripts/localhost.* /home/scripts/
WORKDIR /home
EXPOSE 4003
EXPOSE 5003

ENTRYPOINT [ "./compression-service" ]
