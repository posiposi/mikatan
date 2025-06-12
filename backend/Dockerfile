FROM golang:1.24.4

WORKDIR /go/src

RUN apt-get update && apt-get install -y vim

COPY go.mod go.sum ./

RUN go mod download
RUN go get github.com/steebchen/prisma-client-go
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u gorm.io/gorm
RUN go install github.com/air-verse/air@latest

# Alias for the prisma client
RUN echo 'alias prisma="go run github.com/steebchen/prisma-client-go"' >> ~/.bashrc

COPY . /go/src

CMD ["air", "-c", ".air.toml"]
