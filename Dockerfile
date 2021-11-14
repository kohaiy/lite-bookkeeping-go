FROM golang

WORKDIR /serve

COPY . .

RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct && go build .

EXPOSE 9000

CMD [ "go", "run", "main.go" ]
