FROM alpine:latest

RUN mkdir /app
WORKDIR /app

COPY backApp /app

CMD [ "/app/backApp"]