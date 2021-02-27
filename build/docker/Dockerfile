FROM golang:1.15

WORKDIR /go/src/app

COPY . .

RUN make test
RUN make install

ENTRYPOINT /go/bin/apiserver -log-level $APP_LOG_LEVEL -bind-addr $APP_BIND_ADDR

EXPOSE 8080
