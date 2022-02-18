FROM golang:1.17.7
LABEL maintainer=jupemara
RUN apt-get update --yes && apt-get install --yes \
  ca-certificates
WORKDIR /go/src/github.com/jupemara
COPY . .
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -o /opt/cloudsql-proxy ./main.go
ENTRYPOINT ["/opt/cloudsql-proxy"]
