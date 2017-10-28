# Use the official go docker image built on debian.
#FROM golang:1.7.1
FROM golang:1.8

WORKDIR /go/src/app
COPY . .

# Install revel and the revel CLI.
RUN go get github.com/revel/revel
RUN go get github.com/revel/cmd/revel

# Use the revel CLI to start up our application.
ENTRYPOINT revel run app prod

# Open up the port where the app is running.
EXPOSE 8080
