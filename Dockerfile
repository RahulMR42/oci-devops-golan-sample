FROM golang

WORKDIR /app

COPY * ./


EXPOSE 8080

CMD [ "/sample-program" ]