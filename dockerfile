FROM golang:latest
WORKDIR /usr/src/app
COPY . .
RUN go build -o install-aws-cli
CMD ["./install-aws-cli"]