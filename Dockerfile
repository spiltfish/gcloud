FROM golang
RUN mkdir /mc-worker
ADD . /mc-worker
WORKDIR /mc-worker
RUN go get 	github.com/husobee/vestigo &&
    go get	encoding/json
RUN go build -o main .
CMD ["/mc-worker/main"]