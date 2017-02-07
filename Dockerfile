FROM alpine:latest

RUN echo "Asia/shanghai" >> /etc/timezone

COPY ./main /bin/kk-feedback

RUN chmod +x /bin/kk-feedback

COPY ./config /config

COPY ./app.ini /app.ini

ENV KK_ENV_CONFIG /config/env.ini

VOLUME /config

CMD kk-feedback $KK_ENV_CONFIG

