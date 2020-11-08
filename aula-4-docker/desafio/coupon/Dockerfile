FROM golang:1.15

WORKDIR /coupon/

COPY . .

RUN GOOS=linux go build

EXPOSE 9092

ENTRYPOINT [ "./coupon-microservice" ]