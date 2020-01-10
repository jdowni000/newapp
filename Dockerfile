FROM golang

ADD . /app

WORKDIR /app

RUN go build -v -o newapp

ENTRYPOINT /app/newapp













