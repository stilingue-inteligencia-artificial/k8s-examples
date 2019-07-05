# Start by building the application.
FROM golang:1.12

# setting workdir
WORKDIR /go/src/app
# copping source
COPY *.go .

# Instaling dependencies and building go app
RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080
CMD ["app"]
