FROM golang:1.21.4-alpine3.18

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o server cmd/job-portal-api/main.go

CMD [ "./server" ]