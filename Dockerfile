FROM golang
RUN mkdir /mc-worker
ADD . /mc-worker
WORKDIR /mc-worker
RUN go get 	github.com/husobee/vestigo && go get encoding/json && google.golang.org/api/compute/v1 && golang.org/x/oauth2/google && golang.org/x/net/context
RUN go build -o main .
CMD ["/mc-worker/main"]