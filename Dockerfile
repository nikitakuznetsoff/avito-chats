FROM golang:latest

WORKDIR /go/src/github.com/nikitakuznetsoff/avito-chats
COPY . /go/src/github.com/nikitakuznetsoff/avito-chats/

RUN go build -o ./bin/chatssapp ./cmd/chatssapp/

CMD [ "/go/src/github.com/nikitakuznetsoff/ozon-links-app/bin/chatssapp" ]
