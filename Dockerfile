FROM golang:latest

LABEL manteiner="Salomé Gené <digsach2003@gmail.com>"

WORKDIR /app

COPY go.mod .
COPY go.sum .
 
RUN go mod download

COPY . .

RUN go build

CMD ["./shopping-list"]