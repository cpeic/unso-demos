#First, compile golang app in a build environment
FROM golang:1.16 as build-env

WORKDIR /app
ADD ./code /app/.

#compile go app
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/http-get-app .

#install public certificate authorities
RUN apt update && apt install -y ca-certificates


#Second, copy the compiled binary to a scratch file system
FROM scratch

#copy app
COPY --from=build-env /app/http-get-app .

#copy ssl certificates
COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["./http-get-app"]
