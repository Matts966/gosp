FROM alpine:3.9

ENV GOPATH /go
ENV PATH $PATH:$GOPATH/bin

RUN apk add su-exec dumb-init git go libc-dev
RUN addgroup gosp && \
    adduser -S -G gosp gosp
RUN mkdir $GOPATH \
  && export PATH=$PATH:$GOPATH/bin \
  && go version \
  && go get github.com/Matts966/gosp/cmd/gosp

COPY docker-entrypoint.sh /usr/local/bin/docker-entrypoint.sh
RUN chmod +x /usr/local/bin/docker-entrypoint.sh
ENTRYPOINT ["docker-entrypoint.sh"]
