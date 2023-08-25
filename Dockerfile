FROM golang

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /mytask cmd/main.go

EXPOSE 8080

CMD [ "/mytask" ]