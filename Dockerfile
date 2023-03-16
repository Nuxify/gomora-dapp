FROM golang:1.20

ENV GO111MODULE=on

WORKDIR /app/build
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

CMD ["make", "run"]