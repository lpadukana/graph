FROM alpine:3.3

RUN apk -vv --update --no-cache \
  add font-misc-misc graphviz \
  && rm -rf /var/cache/apk/*

COPY ./graph /opt/bin/graph

CMD [ "/opt/bin/graph" ]
