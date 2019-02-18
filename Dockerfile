FROM golang
ADD . /go/src/github.com/apmath-web/credit-go
WORKDIR /go/src
RUN go get -v github.com/franela/goblin
RUN mkdir build
RUN go build -i -o ./build/server ./github.com/apmath-web/credit-go/application.go
CMD ["./build/server"]