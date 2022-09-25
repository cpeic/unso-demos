#First, compile golang app in a build environment
FROM golang:1.16 as build-env

WORKDIR /app
ADD ./code /app/.

#compile go app
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/http-get-app .

#Second, copy the compiled binary to a scratch file system
FROM scratch

#copy app
COPY --from=build-env /app/http-get-app .

ENTRYPOINT ["./http-get-app"]
