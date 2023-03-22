FROM alpine:latest

RUN mkdir /app
WORKDIR /app

COPY frontApp /app
COPY index.html /app

EXPOSE 8080
CMD [ "/app/frontApp"]