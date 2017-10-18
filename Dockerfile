FROM alpine:latest

COPY ./bin/linux_amd64/kubeapi /usr/local/bin/

ENV APP_BIND_HOST 0.0.0.0

ENV APP_BIND_PORT 1690

CMD [ "kubeapi" ]
