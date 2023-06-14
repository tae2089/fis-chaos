FROM golang:1.19.3-alpine as build
RUN apk --no-cache add tzdata ca-certificates
WORKDIR /src/
COPY . ./
RUN CGO_ENABLED=0 go build -o /bin/main run/main.go
FROM scratch
COPY --from=build /bin/main /bin/main
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENV TZ=Asia/Seoul
ENTRYPOINT ["/bin/main"]