FROM golang:1.19

workdir /app

ENTRYPOINT ["tail", "-f", "/dev/null"]