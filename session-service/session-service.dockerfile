FROM alpine:3.14

RUN mkdir /app

COPY sessionApp /app

CMD [ "/app/sessionApp" ]