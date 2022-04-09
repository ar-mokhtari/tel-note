FROM golang:1.15.3-alpine3.12

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -v -o main .

CMD /*/app/main*/