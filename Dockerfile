FROM alpine:3.20.3

WORKDIR /app

COPY linkoln ./

EXPOSE 8080

CMD [ "/app/linkoln" ]