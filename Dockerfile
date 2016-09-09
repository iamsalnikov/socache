FROM golang:1.7.1

COPY . /go/src/github.com/iamsalnikov/socache
RUN cd /go/src/github.com/iamsalnikov/socache && \
    go get -d -v && \
    go install -v

WORKDIR "/go/src/github.com/iamsalnikov/socache"

ENTRYPOINT socache
EXPOSE 9099
