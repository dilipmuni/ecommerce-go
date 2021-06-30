FROM golang:1.12.0-alpine3.9
RUN apk add git
RUN apk add mongodb
WORKDIR /goapp
ADD . .
## Add this go mod download command to pull in any dependencies
RUN go mod download
ADD . .
## Our project will now successfully build with the necessary go libraries included.
RUN go build -o main .
## our newly created binary executable
CMD ["/goapp/main"]
EXPOSE 9012
