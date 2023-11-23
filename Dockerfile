FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

COPY cmd .

RUN go build -o /delayAlert-order-management-system/cmd

CMD ["/delayAlert-order-management-system/cmd"]