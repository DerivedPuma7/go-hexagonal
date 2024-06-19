FROM golang:1.19

WORKDIR /go/src

# Configure Go module
COPY go.mod .
COPY go.sum .
RUN go mod download

#Install executable
RUN go install github.com/golang/mock/mockgen@v1.5.0
RUN go install github.com/spf13/cobra-cli@latest
RUN ln -s /go/bin/cobra-cli /go/bin/cobra

RUN apt-get update && apt-get install sqlite3 -y

COPY . .

RUN usermod -u 1000 www-data
RUN mkdir -p /var/www/.cache
RUN chown -R www-data:www-data /go
RUN chown -R www-data:www-data /var/www/.cache
USER www-data

CMD ["tail", "-f", "/dev/null"]
