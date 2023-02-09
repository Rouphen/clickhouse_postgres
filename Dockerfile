FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/ecommerce/apigateway
COPY . $GOPATH/src/ecommerce/apigateway
RUN go build -o jenkins_test ./main.go

EXPOSE 30001
ENTRYPOINT ["./jenkins_test"]
