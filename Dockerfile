FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . ./

RUN go build -tags netgo -ldflags "-s -w" -o app-build

EXPOSE 8090

CMD [ "./app-build" ]