FROM  golang:1.22.1-alpine

COPY . /go/src/github.com/nanoandrew4/ngcplogs
RUN cd /go/src/github.com/nanoandrew4/ngcplogs && go get && go build --ldflags '-extldflags "-static"' -o /usr/bin/ngcplogs

WORKDIR /usr/bin
ENTRYPOINT [ "/usr/bin/ngcplogs" ]