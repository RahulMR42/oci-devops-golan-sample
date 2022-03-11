FROM golang

WORKDIR /app

COPY * .

RUN go build -o main .

EXPOSE 8080

CMD [ "sample-program" ]