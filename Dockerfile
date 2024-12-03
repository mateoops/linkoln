FROM alpine:3.20.3

WORKDIR /app

COPY linkoln ./

RUN addgroup -S appgroup && adduser -S appuser -G appgroup \
    && chown -R appuser:appgroup /app

USER appuser

EXPOSE 8080

CMD [ "/app/linkoln" ]