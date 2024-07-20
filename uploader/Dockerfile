FROM golang:1.22.5

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o apostoli-uploader

CMD ["./apostoli-uploader"]

