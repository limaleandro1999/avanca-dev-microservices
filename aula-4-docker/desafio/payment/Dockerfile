FROM golang:1.15

WORKDIR /payment/

COPY . .

RUN GOOS=linux go build

EXPOSE 9091

ENTRYPOINT [ "./payment-microservice" ]