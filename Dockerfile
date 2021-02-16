FROM golang:1.15-alpine3.12
RUN apk add git
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/gorilla/mux

WORKDIR /app
COPY . ./
EXPOSE 8081
CMD ["go", "run", "main.go"]
