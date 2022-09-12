#First, compile golang app in a build environment
FROM golang:1.16 as build-env

WORKDIR /app
ADD ./code /app/.

RUN GOOS=linux go build -o /app/hello-world-app .


#Second, copy the compiled binary to a scratch file system
FROM scratch

COPY --from=build-env /app/hello-world-app .

ENTRYPOINT ["./hello-world-app"]
