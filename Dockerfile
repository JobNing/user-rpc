FROM golang:1.22.10

RUN mkdir /app
#相当于linux跑的cd
WORKDIR /app

COPY ./ ./

RUN go build main.go

CMD ["./main"]

