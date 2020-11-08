FROM golang:1.15

WORKDIR /checkout/

COPY . .

RUN GOOS=linux go build

EXPOSE 9090

ENTRYPOINT [ "./checkout-microservice" ]