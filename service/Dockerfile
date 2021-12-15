FROM golang:latest

WORKDIR /app
COPY service.go ./app/service.go
RUN cd ./app \
    && go mod init main.com/main \
    && go get github.com/go-sql-driver/mysql \
    && go get github.com/gin-gonic/gin \
    && go get . \
    && go build service.go 

ENTRYPOINT [ "/app/service" ]
EXPOSE 3000
