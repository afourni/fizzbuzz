FROM scratch

LABEL Description="This image is used to start the FizzBuzz API" Version="1.0"

ENV GIN_MODE=release

EXPOSE 8080

ADD fizzbuzz /

CMD ["/fizzbuzz"]
